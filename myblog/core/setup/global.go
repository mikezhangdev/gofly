package setup

import (
	"fmt"
	"log"
	"myblog/core"
	"myblog/core/setting"
	//"myblog/core/app"
)

func SetGlobalLogger(app *core.Application) error {
	l, err := NewZapLogger(&ZapLoggerCfg{
		DevStdOut: true, // 通过识别环境变量 获取开发环境标识 TODO
		AppName:   app.Name,
	})
	if err != nil {
		return err
	}
	core.Logger = l
	return nil
}

func NewGlobalVars(application *core.Application, cfg *setting.Config) error {
	// 初始化Mysql数据库
	if cfg.MysqlCfg.Host != "" {
		mysqlConn, err := NewMysqlConn(&MysqlCfg{
			User:      cfg.MysqlCfg.User,
			Pass:      cfg.MysqlCfg.Pass,
			Host:      cfg.MysqlCfg.Host,
			Port:      cfg.MysqlCfg.Port,
			DbName:    cfg.MysqlCfg.DbName,
			Charset:   cfg.MysqlCfg.Charset,
			ParseTime: cfg.MysqlCfg.ParseTime,
			Loc:       cfg.MysqlCfg.Loc,
		})
		if err != nil {
			return fmt.Errorf("NewMysqlConn err %v", err)
		}
		core.MysqlConn = mysqlConn
	}

	log.Printf("NewRedis-0000000000000000")
	if cfg.RedisCfg.Addr != ""{
		log.Printf("NewRedis-1111111111")
		redisP,err := NewRedis(&RedisCfg{
			Addr:            cfg.RedisCfg.Addr,
			//MaxIdle:         cfg,
			MaxActive:       0,
			IdleTimeout:     0,
			MaxConnLifetime: 0,
		})
		if err != nil{
			return err
		}
		core.RedisPool = redisP
	}

	return nil
}
