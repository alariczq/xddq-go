package assist

import "go.uber.org/zap"

func (a *Assistant) L() *zap.Logger {
	return a.logger.Load().(*zap.Logger)
}

func (a *Assistant) SetLogger(logger *zap.Logger) {
	a.c.SetLogger(logger)
	a.logger.Store(logger)
}
