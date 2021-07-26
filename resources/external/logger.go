package external

import (
	"github.com/ravielze/Labpro-BE/config"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/logs/zap"
	z "go.uber.org/zap"
)

func NewLogger(config *config.Env) (logs.Logger, error) {
	return zap.New(config.IsDevelopment(), zap.Option{
		Level:  logs.GetLoggerLevel(config.LogLevel),
		Prefix: "",
	}, z.AddStacktrace(z.ErrorLevel), z.AddCallerSkip(1))
}
