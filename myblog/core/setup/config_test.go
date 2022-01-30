package setup

import (
	"github.com/spf13/viper"
	"testing"
)

type Config struct {
	System   System `mapstructure:"system" json:"system" yaml:"system"`
	MysqlCfg Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	RedisCfg Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	// `mapstructure:"System" json:"System" yaml:"System"`
}

//system:
//dev_std_out: "8081"
//
//mysql:
//user: "root"
//pass: "123456"
//host: "192.168.56.104"
//port: 3306
//db_name: "myblog"
//charset: "utf8mb4"
//parse_time: "true"
//loc: "Local"
//
//redis:
//addr: "192.168.56.104:6379"
//max_idle: 20
//max_active: 30
//idle_timeout: 10
//max_conn_lifetime: 60

type System struct {
	DevStdOut string `yaml:"dev_std_out" json:"dev_std_out"`
}

type Mysql struct {
	User      string `yaml:"user" json:"user"`
	Pass      string `yaml:"pass" json:"pass"`
	Host      string `yaml:"host" json:"host"`
	Port      int64  `yaml:"port" json:"port"`
	DbName    string `yaml:"dbname" json:"dbname"`
	Charset   string `yaml:"charset" json:"charset"`
	ParseTime string `yaml:"parsetime" json:"parsetime"`
	Loc       string `yaml:"loc" json:"loc"`
}

type Redis struct {
	Addr        string `json:"addr" yaml:"addr"`
	MaxIdle     int    `yaml:"maxidle" json:"maxidle"` // 具体配置参数 根据具体情况配置 待优化研究细化 TODO
	MaxActive   int    `yaml:"maxactive" json:"maxactive"`
	IdleTimeout int    `yaml:"idletimeout" json:"idletimeout"`
	//Wait            bool
	MaxConnLifetime int `yaml:"maxconnlifetime" json:"maxconnlifetime"`
}


func TestNewViperConfig(t *testing.T) {
	var config Config
	_, err := NewViperConfig(&ViperConfig{
		ConfigPath: "./",
		ConfigName: "config_test",
		LoadInnerConfig: func(v *viper.Viper) error {
			err := v.Unmarshal(&config)
			if err != nil {
				return nil
			}
			return nil
		},
	})
	if err != nil {
		t.Fatalf("NewViperConfig err %v", err)
	}
	//log.Printf("%v",config)
	//if config.System.DevStdOut != "dev" {
	//	t.Fatalf("yaml parse not right")
	//}
}
