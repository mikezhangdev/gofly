package main

import (
	"go.uber.org/zap"
	"log"
	"myblog/middleware/zaplog"
	"time"
)

func main() {

	zaplog, err := zaplog.NewZapLogger(&zaplog.ZapLoggerCfg{
		DevStdOut: true,
		AppName:   "myblog",
	})
	if err != nil {
		log.Fatalf("NewZapLogger err %v", err)
		return
	}
	zaplog.Info("333", zap.String("time", time.Now().Format("2006-01-01 15:04:05")))
	zaplog.Info("333444", zap.String("time", time.Now().Format("2006-01-01 15:04:05")))
	zaplog.Info("333", zap.String("time", time.Now().Format("2006-01-01 15:04:05")))
	zaplog.Info("333444", zap.String("time", time.Now().Format("2006-01-01 15:04:05")))
	//zaplog.Sync()

}
