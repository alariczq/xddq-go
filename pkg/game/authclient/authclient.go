package authclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"slices"

	"xddq/pkg/logutil"
	"xddq/pkg/utils"

	"github.com/imroc/req/v3"
)

var defaultAuthClient = NewAuthClient()

type AuthClient struct {
	cli *req.Client
}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"

func NewAuthClient() *AuthClient {
	cli := req.C().
		SetUserAgent(userAgent)

	if os.Getenv("DEBUG") != "" {
		cli = cli.DevMode()
	}

	return &AuthClient{
		cli: cli,
	}
}

type ServerInfo struct {
	ServerId   int64  `json:"serverId"`
	LabelName  string `json:"labelName"`
	ServerName string `json:"serverName"`
	Address    string `json:"address"`
}

type ListServersResult struct {
	ServerList        []ServerInfo `json:"serverList"`
	PlayerServerList  []int64      `json:"playerServerList"`
	RecommendServerId int64        `json:"recommendServerId"`
}

func (c *AuthClient) ListServers(ctx context.Context, username, password string) (result *ListServersResult, err error) {
	lr, err := c.login(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return c.ListServersWithUid(ctx, lr.Userinfo.Uid)
}

func (c *AuthClient) ListServersWithUid(ctx context.Context, uid string) (result *ListServersResult, err error) {
	_, err = c.cli.R().
		SetContext(ctx).
		SetBodyJsonMarshal(map[string]any{
			"openId":    uid,
			"channelId": 31,
		}).
		SetSuccessResult(&result).
		Post("https://login-xddq.hdnd01.com/server/list")
	return
}

func encryptPwd(pwd string) string {
	randomStr := utils.RandomString(8) + pwd[:3] + utils.RandomString(5) + pwd[3:] + utils.RandomString(2)
	return base64.StdEncoding.EncodeToString([]byte(randomStr))
}

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type loginRequest struct {
	LoginAccount string `json:"login_account"`
	Password     string `json:"password"`
}

type loginResult struct {
	AppPst   string `json:"app_pst"`
	Userinfo struct {
		Uid string `json:"uid"`
	} `json:"userinfo"`
}

func (c *AuthClient) login(ctx context.Context, username, password string) (result *loginResult, err error) {
	resp := &Response{Data: &result}
	_, err = c.cli.R().
		SetContext(ctx).
		SetBodyJsonMarshal(loginRequest{
			LoginAccount: username,
			Password:     encryptPwd(password),
		}).
		SetSuccessResult(resp).
		Post("https://mysdk.37.com/index.php?c=api-login&a=act_login")
	if err != nil {
		return nil, err
	}

	if resp.Code != 1 {
		return nil, fmt.Errorf("login failed: %d:%s", resp.Code, resp.Msg)
	}

	return
}

type apiLoginResult struct {
	AppPst string `json:"app_pst"`
}

func (c *AuthClient) apiLogin(username, ptoken string) (result *apiLoginResult, err error) {
	resp := &Response{Data: &result}
	_, err = c.cli.R().
		SetQueryParams(map[string]string{
			"is_self": "1",
			"pid":     "37h5",
			"ptoken":  ptoken,
			"puid":    username,
		}).
		SetSuccessResult(resp).
		Post("https://apimyh5.37.com/index.php?c=sdk-login&a=act_login")
	if err != nil {
		return nil, err
	}

	if resp.Code != 1 {
		return nil, errors.New(resp.Msg)
	}

	return
}

type RequestData struct {
	ClientId string `json:"clientId"`
	Token    string `json:"token"`
	Uid      string `json:"uid"`
	Uname    string `json:"uname"`
}

func (r *RequestData) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(*r)
	if err != nil {
		return nil, err
	}

	return []byte(`"` + url.QueryEscape(string(b)) + `"`), nil
}

type LoginRequest struct {
	Data      *RequestData `json:"data"`
	LoginType int32        `json:"loginType"`
	ChannelId int32        `json:"channelId"`
	Appid     string       `json:"appid"`
	GameId    int32        `json:"gameId"`
}

type LoginResult struct {
	Ret          int32  `json:"ret"`
	PlayerId     int64  `json:"playerId"`
	OpenId       string `json:"openId"`
	NickName     string `json:"nickName"`
	Token        string `json:"token"`
	WsAddress    string `json:"wsAddress"`
	CreatePlayer bool   `json:"createPlayer"`
	RealmsId     int32  `json:"realmsId"`

	ServerId int64  `json:"serverId"`
	Uid      string `json:"uid"`
	AppPst   string `json:"appPst"`
}

type LoginOptions struct {
	ServerId int64
	Uid      string
	appPst   string
	Address  string
}

type LoginOption func(opts *LoginOptions)

func WithServerId(serverId int64) LoginOption {
	return func(opts *LoginOptions) {
		opts.ServerId = serverId
	}
}

func WithUid(uid string) LoginOption {
	return func(opts *LoginOptions) {
		opts.Uid = uid
	}
}

func WithAppPst(pst string) LoginOption {
	return func(opts *LoginOptions) {
		opts.appPst = pst
	}
}

func (c *AuthClient) Login(ctx context.Context, username string, password string, opts ...LoginOption) (result *LoginResult, err error) {
	options := &LoginOptions{}
	for _, opt := range opts {
		opt(options)
	}

	if options.Uid == "" || options.appPst == "" {
		logutil.Info(ctx, "login with username and password")
		lr, err := c.login(ctx, username, password)
		if err != nil {
			return nil, err
		}

		options.Uid = lr.Userinfo.Uid
		options.appPst = lr.AppPst
	}

	if options.ServerId == 0 || options.Address == "" {
		result, err := c.ListServersWithUid(ctx, options.Uid)
		if err != nil {
			return nil, err
		}

		if options.ServerId == 0 {
			options.ServerId = result.RecommendServerId
		}

		idx := slices.IndexFunc(result.ServerList, func(s ServerInfo) bool {
			return s.ServerId == options.ServerId
		})
		if idx == -1 {
			return nil, fmt.Errorf("server %d not found", options.ServerId)
		}
		options.Address = result.ServerList[idx].Address
	}

	_, err = c.cli.R().
		SetContext(ctx).
		SetBodyJsonMarshal(LoginRequest{
			Data: &RequestData{
				ClientId: utils.RandomString(16),
				Token:    options.appPst,
				Uid:      options.Uid,
				Uname:    username,
			},
			LoginType: 0,
			ChannelId: 31,
			Appid:     "37h5",
			GameId:    223,
		}).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("%s/player/login", options.Address))
	if err != nil {
		return nil, err
	}

	if result.Ret != 0 {
		return nil, fmt.Errorf("login failed: %d", result.Ret)
	}

	result.Uid = options.Uid
	result.ServerId = options.ServerId
	result.AppPst = options.appPst

	return
}

func Login(ctx context.Context, username, password string, opts ...LoginOption) (result *LoginResult, err error) {
	return defaultAuthClient.Login(ctx, username, password, opts...)
}
