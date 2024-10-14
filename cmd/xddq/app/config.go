package app

import (
	"errors"
	"xddq/pkg/assist"
)

type Config struct {
	Addr     string
	LogLevel string
	Accounts []*assist.AssistantConfig
}

func (c *Config) Validate() error {
	if c.Addr == "" {
		return errors.New("addr is required")
	}

	if c.LogLevel != "debug" && c.LogLevel != "info" && c.LogLevel != "warn" && c.LogLevel != "error" {
		return errors.New("log level must be debug, info, warn or error")
	}

	for _, account := range c.Accounts {
		if err := account.Validate(); err != nil {
			return err
		}
	}

	return nil
}
