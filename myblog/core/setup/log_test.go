package setup

import (
	"context"
	"testing"
	//"time"
)

func TestNewZapLogger(t *testing.T) {
	zaplog, err := NewZapLogger(&ZapLoggerCfg{
		DevStdOut: true,
		AppName:   "myblog",
	})
	if err != nil {
		t.Fatalf("NewZapLogger err %v", err)
	}
	zaplog.InfoF(context.Background(),"333")
}
