package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"myblog/core/logging"
)

type Application struct {
	Name       string
	Addr       string
	ConfigPath string
	ConfigName string
	LoadConfig func(v *viper.Viper) error
	GinHandler func(r *gin.Engine) error
}

var Logger logging.LoggerIFace

var MysqlConn *gorm.DB


var RedisPool *redis.Pool
