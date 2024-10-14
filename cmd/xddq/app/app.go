package app

import (
	"time"
	"xddq/pkg/assist"
	"xddq/pkg/logutil"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func initLogger(conf *Config) error {
	logger, err := logutil.NewZapLogger(conf.LogLevel)
	if err != nil {
		return err
	}
	logger = logger.WithOptions(zap.WithCaller(false))

	logutil.SetDefaultZapLogger(logger)

	return nil
}

type App struct {
	conf    *Config
	server  *server.Hertz
	assists map[string]*assist.Assistant
}

func NewApp(conf *Config) *App {
	app := &App{
		conf:    conf,
		assists: make(map[string]*assist.Assistant),
	}

	hlog.SetLevel(hlog.LevelError)

	app.server = server.New(
		server.WithHostPorts(conf.Addr),
		server.WithDisablePrintRoute(true),
	)

	return app
}

func (a *App) Run() error {
	a.initServer()
	a.startAssists()

	return a.server.Run()
}

func (a *App) startAssists() {
	for _, conf := range a.conf.Accounts {
		if conf.Disabled {
			continue
		}

		logutil.DefaultLogger.Debug("start assist", zap.String("username", conf.Username))

		assist := assist.NewAssistant(conf)

		a.assists[conf.Username] = assist

		go assist.Run()

		time.Sleep(10 * time.Second)
	}
}

func (a *App) stopAssists() {
	for _, assist := range a.assists {
		assist.Stop()
	}
}

func (a *App) restartAssists() {
	a.stopAssists()
	a.startAssists()
}

func Main() error {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	conf := &Config{}
	if err := viper.Unmarshal(conf); err != nil {
		return err
	}

	if err := conf.Validate(); err != nil {
		return err
	}

	initLogger(conf)

	app := NewApp(conf)

	if err := app.Run(); err != nil {
		return err
	}

	return nil
}
