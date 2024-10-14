package client

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"xddq/pkg/game/authclient"
	"xddq/pkg/game/db"
	"xddq/pkg/game/protocol"
	"xddq/pkg/game/protocol/pb"
	"xddq/pkg/logutil"
	"xddq/pkg/utils"

	"github.com/coder/websocket"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	o *ClientOptions

	conn *websocket.Conn

	sc      chan *message
	running atomic.Bool
	closeMu sync.Mutex
	logined atomic.Bool

	ctx    context.Context
	cancel context.CancelFunc

	limiter ratelimit.Limiter

	errHandler func(p *protocol.Packet, err error)

	onLogined      func(Context) error
	mu             sync.RWMutex
	defaultHandler Handler
	ignoreCmds     map[int32]bool
	handlers       map[int32]Handler
	logger         atomic.Value

	lastPingTime   atomic.Value
	networkLatency time.Duration
}

func NewClient(username, password string, opts ...ClientOption) *Client {
	o := &ClientOptions{
		username: username,
		password: password,
		ctx:      context.Background(),
		timeout:  10 * time.Second,
	}
	for _, opt := range opts {
		opt(o)
	}

	c := &Client{
		o:        o,
		sc:       make(chan *message, 10),
		limiter:  ratelimit.New(20),
		handlers: make(map[int32]Handler),
		ignoreCmds: map[int32]bool{
			protocol.PLAYER_PING: true,
		},
	}

	c.logger.Store(logutil.DefaultLogger)

	c.OnReqLoginMsg(c.onLoginRsp)
	c.OnReqPingMsg(c.onPingRsp)

	return c
}

func (c *Client) SetLogger(logger *zap.Logger) {
	c.logger.Store(logger)
}

func (c *Client) l() *zap.Logger {
	return c.logger.Load().(*zap.Logger)
}

func (c *Client) connect(ctx context.Context) error {
	if c.o.username == "" || c.o.password == "" {
		return errors.New("username or password is empty")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := authclient.Login(ctx, c.o.username, c.o.password,
		authclient.WithAppPst(c.o.token),
		authclient.WithServerId(c.o.serverId),
		authclient.WithUid(c.o.uid),
	)
	if err != nil {
		return err
	}

	c.o.token = result.AppPst
	c.o.uid = result.Uid
	c.o.serverId = result.ServerId
	c.o.playerId = result.PlayerId

	conn, _, err := websocket.Dial(ctx, result.WsAddress, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		return err
	}

	conn.SetReadLimit(5 * 1024 * 1024)

	c.conn = conn

	return c.login(ctx, result.Token)
}

func (c *Client) onLoginRsp(ctx Context, msg *pb.RspLoginMsg) error {
	if msg.GetRet() != 0 {
		return db.GetGameError(msg.GetRet())
	}

	c.logined.Store(true)

	if c.onLogined != nil {
		return c.onLogined(ctx)
	}

	return nil
}

func (c *Client) Close() error {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()

	if !c.running.Load() {
		return nil
	}

	c.running.Store(false)

	c.cancel()

	return c.conn.CloseNow()
}

func (c *Client) ServerId() int64 {
	return c.o.serverId
}

func (c *Client) PlayerId() int64 {
	return c.o.playerId
}

func (c *Client) send(ctx context.Context, msgId int32, data []byte) (err error) {
	m := &message{
		packet: protocol.NewPacket(msgId, data).WithPlayerId(c.o.playerId),
		errC:   make(chan error, 1),
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c.ctx.Done():
		return c.ctx.Err()
	case c.sc <- m:
		if !c.running.Load() {
			return nil
		}
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c.ctx.Done():
		return c.ctx.Err()
	case err := <-m.errC:
		return err
	}
}

func (c *Client) Send(ctx context.Context, msgId int32, data []byte) (err error) {
	return c.send(ctx, msgId, data)
}

func (c *Client) SendMessage(ctx context.Context, msg pb.Message) (err error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return c.send(ctx, msg.MsgId(), data)
}

func (c *Client) WriteMessage(ctx context.Context, mt websocket.MessageType, data []byte) (err error) {
	return c.conn.Write(ctx, mt, data)
}

func (c *Client) ReadMessage(ctx context.Context) (mt websocket.MessageType, data []byte, err error) {
	return c.conn.Read(ctx)
}

func (c *Client) RegisterHandler(msgId int32, handler Handler) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.handlers[msgId]; ok {
		if cmd := db.GetRecvCmd(msgId); cmd != nil {
			return fmt.Errorf("handler already registered for msgId %d, cmd %s", msgId, cmd.SmMethod)
		}

		return fmt.Errorf("handler already registered for msgId %d", msgId)
	}

	c.handlers[msgId] = handler
	return nil
}

func (c *Client) WithDefaultHandler(handler Handler) {
	c.defaultHandler = handler
}

func (c *Client) WithOnLogined(fn func(Context) error) {
	c.onLogined = fn
}

type message struct {
	packet *protocol.Packet
	errC   chan error
}

func (c *Client) login(ctx context.Context, token string) error {
	return c.SendReqLoginMsg(ctx, &pb.ReqLoginMsg{
		Token:        &token,
		Language:     utils.Ptr("zh_cn"),
		LiveShowType: utils.Ptr(int32(1)),
	})
}

func (c *Client) ping(ctx context.Context, wg *sync.WaitGroup) error {
	wg.Done()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-ticker.C:
			c.lastPingTime.Store(time.Now())
			if err := c.writeMessage(ctx, protocol.NewPacket(protocol.PLAYER_PING, nil)); err != nil {
				return err
			}
		}
	}
}

func (c *Client) onPingRsp(_ Context, _ *pb.RspPingMsg) error {
	v, _ := c.lastPingTime.Load().(time.Time)
	c.networkLatency = time.Since(v)
	c.l().Info("🛜 Networkd Latency", zap.Duration("networkLatency", c.networkLatency))
	return nil
}

func (c *Client) writeMessage(ctx context.Context, p *protocol.Packet) error {
	l := c.l().With(
		zap.Int32("msgId", p.MsgId),
		zap.Int("size", len(p.Body)),
	)

	cmd := db.GetCmd(p.MsgId)
	if cmd == nil {
		l.Debug("🚀 Unknown message")
	}

	if cmd != nil && !c.ignoreCmds[cmd.CmMsgId] {
		l.Debug(fmt.Sprintf("🚀 %s", cmd.Comment), zap.String("cmd", cmd.CmMethod))
	}

	return c.conn.Write(ctx, websocket.MessageBinary, p.Encode())
}

func (c *Client) sendLoop(ctx context.Context, wg *sync.WaitGroup) error {
	wg.Done()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case msg, ok := <-c.sc:
			if !ok {
				return errors.New("send loop closed")
			}

			c.limiter.Take()

			err := c.writeMessage(ctx, msg.packet)

			msg.errC <- err

			if err != nil {
				return err
			}

		}
	}
}

func (c *Client) readLoop(ctx context.Context, wg *sync.WaitGroup) error {
	wg.Done()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			_, r, err := c.conn.Reader(ctx)
			if err != nil {
				return errors.New("connection closed")
			}

			msg, err := protocol.Decode(r)
			if err != nil {
				c.l().Error("decode message failed", zap.Error(err))
				continue
			}

			c.handlePacket(msg)
		}
	}
}

func (c *Client) handlePacket(p *protocol.Packet) {
	l := c.l().With(
		zap.Int32("msgId", p.MsgId),
		zap.Int("size", len(p.Body)),
	)

	cmd := db.GetRecvCmd(p.MsgId)
	if cmd == nil {
		l.Debug("⬇️ Unknown message")
	}

	if cmd != nil && !c.ignoreCmds[cmd.CmMsgId] {
		l.Debug(fmt.Sprintf("⬇️ %s", cmd.Comment), zap.String("cmd", cmd.SmMethod))
		l = l.With(zap.String("cmd", cmd.SmMethod), zap.String("comment", cmd.Comment))
	}

	c.mu.RLock()
	handler, ok := c.handlers[p.MsgId]
	c.mu.RUnlock()
	if !ok {
		if c.defaultHandler == nil {
			return
		}

		handler = c.defaultHandler
	}

	executeHandler := func() {
		ctx, cancel := context.WithTimeout(c.ctx, c.o.timeout)
		defer cancel()
		cc := &clientContext{
			Context: ctx,
			client:  c,
			msg:     p,
		}

		if err := handler(cc); err != nil {
			if c.errHandler != nil {
				c.errHandler(p, err)
				return
			}

			l.Error("handle message failed", zap.Error(err))
		}
	}

	executeHandler()
}

func (c *Client) run() error {
	c.ctx, c.cancel = context.WithCancel(c.o.ctx)

	if err := c.connect(c.ctx); err != nil {
		return err
	}
	defer c.Close()

	eg, ctx := errgroup.WithContext(c.ctx)

	// wait loop running
	wg := &sync.WaitGroup{}
	wg.Add(3)

	eg.Go(func() error {
		return c.sendLoop(ctx, wg)
	})

	eg.Go(func() error {
		return c.readLoop(ctx, wg)
	})

	eg.Go(func() error {
		return c.ping(ctx, wg)
	})

	wg.Wait()

	if c.o.afterConnected != nil {
		c.o.afterConnected(c)
	}

	c.running.Store(true)

	if err := eg.Wait(); err != nil {
		c.l().Error("client run failed", zap.Error(err))
		return err
	}

	return nil
}

func (c *Client) Run() error {
	return c.run()
}
