package app

import (
	"github.com/spf13/viper"
	"log"
	"myblog/core"
	"myblog/core/setting"
	"myblog/core/setup"
)

type Runner struct {
	app    *core.Application
	config *setting.Config
}

func NewRunner(app *core.Application) (*Runner, error) {
	var err error
	// 初始化 配置
	var config setting.Config
	_, err = setup.NewViperConfig(&setup.ViperConfig{
		ConfigPath: app.ConfigPath,
		ConfigName: app.ConfigName,
		LoadConfig: app.LoadConfig,
		LoadInnerConfig: func(v *viper.Viper) error {
			err := v.Unmarshal(&config)
			if err != nil {
				return err
			}
			log.Printf("config %v",config)
			return nil
		},
	})
	if err != nil {
		return nil, err
	}
	// 初始化 日志
	err = setup.SetGlobalLogger(app)
	if err != nil {
		return nil, err
	}
	// 初始化全局变量
	err = setup.NewGlobalVars(app, &config)
	if err != nil {
		return nil, err
	}
	runner := &Runner{
		app:    app,
		config: &config,
	}
	return runner, nil
}
