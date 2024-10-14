package client

import (
	"context"
	"time"
)

type ClientOptions struct {
	username string
	password string

	serverId int64

	uid      string
	token    string
	playerId int64

	ctx            context.Context
	timeout        time.Duration
	afterConnected func(c *Client)
}

type ClientOption func(*ClientOptions)

func WithContext(ctx context.Context) ClientOption {
	return func(c *ClientOptions) {
		c.ctx = ctx
	}
}

func WithServerId(serverId int64) ClientOption {
	return func(c *ClientOptions) {
		c.serverId = serverId
	}
}

func WithToken(token string) ClientOption {
	return func(co *ClientOptions) {
		co.token = token
	}
}

func WithAfterConnected(fn func(c *Client)) ClientOption {
	return func(co *ClientOptions) {
		co.afterConnected = fn
	}
}
