package setup

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlCfg struct {
	User      string
	Pass      string
	Host      string
	Port      int64
	DbName    string
	Charset   string
	ParseTime string
	Loc       string
}

func NewMysqlConn(c *MysqlCfg) (*gorm.DB, error) { // 超时相关配置优化 TODO
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.DbName,
		c.Charset,
		c.ParseTime,
		c.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
