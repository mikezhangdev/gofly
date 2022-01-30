package setup

import (
	"log"
	"testing"
)

func TestNewRedis(t *testing.T) {
	p, err := NewRedis(&RedisCfg{
		Addr:            "192.168.56.104:6379",
		PassWord:        "",
		MaxIdle:         0,
		MaxActive:       0,
		IdleTimeout:     0,
		MaxConnLifetime: 0,
	})
	if err != nil {
		t.Fatalf("NewRedis err %v", err)
	}
	conn := p.Get()
	ret, err := conn.Do("PING")
	if err != nil {
		t.Fatalf("conn.Do err %v", err)
	}
	log.Printf("ret %v", ret)
}
