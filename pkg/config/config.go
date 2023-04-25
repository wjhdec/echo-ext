package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	once sync.Once
	cfg  *Config
)

// New 新建配置，可以设置多个配置读取位置
func New(path ...string) *Config {
	once.Do(func() {
		v := viper.New()
		v.AddConfigPath("./")
		v.AddConfigPath("./configs")
		for _, p := range path {
			v.AddConfigPath(p)
		}
		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}
		cfg = &Config{config: v}
	})
	return cfg
}

type Config struct {
	config *viper.Viper
}

// Reload 重写加载
func (c *Config) Reload() error {
	return c.config.ReadInConfig()
}

func (c *Config) Unmarshal(key string, v any) error {
	return c.config.Sub(key).Unmarshal(v)
}

func (c *Config) StrValue(key string) string {
	return c.config.GetString(key)
}
