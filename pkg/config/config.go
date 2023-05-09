package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	Reload() error
	UnmarshalByKey(key string, v any) error
	ValueByKey(key string) any
	StrValueByKey(key string) string
	ConfigFileUsed() string
	SetByKey(key string, v any) error
}

// New 新建配置，可以设置多个配置读取位置
func New(path ...string) (Config, error) {
	v := viper.New()
	for _, p := range path {
		v.AddConfigPath(p)
	}
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return &config{Viper: *v}, nil
}

type config struct {
	viper.Viper
}

// Reload 重写加载
func (c *config) Reload() error {
	return c.Viper.ReadInConfig()
}

func (c *config) UnmarshalByKey(key string, v any) error {
	return c.Viper.Sub(key).Unmarshal(v)
}

func (c *config) SetByKey(key string, v any) error {
	c.Viper.Set(key, v)
	return nil
}

func (c *config) ValueByKey(key string) any {
	return c.Viper.Get(key)
}

func (c *config) StrValueByKey(key string) string {
	return c.Viper.GetString(key)
}

func (c *config) ConfigFileUsed() string {
	return c.Viper.ConfigFileUsed()
}
