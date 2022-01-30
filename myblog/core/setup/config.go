package setup

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	ConfigPath      string
	ConfigName      string
	LoadConfig      func(v *viper.Viper) error
	LoadInnerConfig func(v *viper.Viper) error
}

func NewViperConfig(c *ViperConfig) (*viper.Viper, error) {
	vp := viper.New()
	vp.AddConfigPath(c.ConfigPath)
	vp.SetConfigName(c.ConfigName)
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("NewViperConfig err %v", err)
	}
	vp.OnConfigChange(func(e fsnotify.Event) { // 监听文件变化动态加载 TODO
		fmt.Printf("vp.OnConfigChange .Name %v", e.Name)
	})
	vp.WatchConfig()
	if c.LoadInnerConfig == nil { // 项目启动需要内置配置
		return nil, fmt.Errorf("c.LoadInnerConfig is nil")
	}
	err = c.LoadInnerConfig(vp)
	if err != nil {
		return nil, fmt.Errorf("c.LoadInnerConfig err %v", err)
	}
	if c.LoadConfig != nil {
		err = c.LoadConfig(vp)
		if err != nil {
			return nil, fmt.Errorf("NewViperConfig err %v", err)
		}
	}

	//vpByte, _ := json.Marshal(*vp)
	fmt.Printf("\nvp %v\n",vp)
	return vp, nil
}
