package setup

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type RedisCfg struct {
	Addr            string
	PassWord        string
	MaxIdle         int  // 具体配置参数 根据具体情况配置 待优化研究细化 TODO
	MaxActive       int
	IdleTimeout     int
	//Wait            bool
	MaxConnLifetime int
}

func NewRedis(c *RedisCfg) (*redis.Pool, error) {
	p := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", c.Addr)
			if err != nil {
				return nil, err
			}
			if c.PassWord != ""{
				_,err = conn.Do("AUTH",c.PassWord)
				if err != nil{
					return nil,fmt.Errorf("redis.Dial err %v",err)
				}
			}
			return conn, nil
		},
		DialContext: nil,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			pong, err := c.Do("PING")
			if err != nil {
				return err
			}
			log.Printf("TestOnBorrow ping:%v", pong)
			return nil
		},
		MaxIdle:         c.MaxIdle,
		MaxActive:       c.MaxActive,
		IdleTimeout:     time.Duration(int64(c.IdleTimeout)) * time.Second,
		//Wait:            c.Wait,
		MaxConnLifetime: time.Duration(int64(c.MaxConnLifetime)) * time.Second,
	}
	return p, nil
}
