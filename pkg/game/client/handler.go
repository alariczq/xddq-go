package client

import (
	"context"

	"xddq/pkg/game/protocol"
	"xddq/pkg/game/protocol/pb"

	"google.golang.org/protobuf/proto"
)

type Handler func(Context) error

type Context interface {
	context.Context
	Client() *Client
	PlayerId() int64
	Message() *protocol.Packet
	Body() []byte
	Send(ctx context.Context, msg pb.Message) error
}

type clientContext struct {
	context.Context
	client *Client
	msg    *protocol.Packet
}

func (c *clientContext) Client() *Client {
	return c.client
}

func (c *clientContext) Message() *protocol.Packet {
	return c.msg
}

func (c *clientContext) Body() []byte {
	return c.msg.Body
}

func (c *clientContext) PlayerId() int64 {
	return c.client.PlayerId()
}

func (c *clientContext) Send(ctx context.Context, msg pb.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	return c.client.Send(ctx, msg.MsgId(), data)
}
