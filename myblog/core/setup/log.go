package setup

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"myblog/core/logging"
	"os"
)

type ZapLoggerCfg struct {
	DevStdOut bool
	AppName   string
}

func NewZapLogger(c *ZapLoggerCfg) (*logging.LoggerContext, error) {
	logDir := "/var/log/service/" + c.AppName + "/"
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	infoLogPath := logDir + "info.log"
	infoWriteSyncer, err := logging.CreateFile(infoLogPath)
	if err != nil {
		return nil, fmt.Errorf("NewZapLogger err %v", err)
	}
	infoLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zap.ErrorLevel
	})
	errLogPath := logDir + "err.log"
	errWriteSyncer, err := logging.CreateFile(errLogPath)
	if err != nil {
		return nil, fmt.Errorf("NewZapLogger err %v", err)
	}
	errLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, infoWriteSyncer, infoLevelEnabler),
		zapcore.NewCore(encoder, errWriteSyncer, errLevelEnabler),
	)
	if c.DevStdOut {
		std := os.Stdout
		//std := zapcore.Lock(os.Stdout)
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(encoder, std, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return true
			})),
		)
	}
	l := zap.New(core)
	loggerContext := logging.NewLoggerContext(l)
	return loggerContext, nil
}
