package main

import "github.com/spf13/viper"

type DemoConfig struct {
	viper.Viper
}

func NewConfig() *DemoConfig {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../../configs/")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return &DemoConfig{*v}
}

func (c *DemoConfig) Reload() error {
	return c.ReadInConfig()
}

// UnmarshalByKey 根据key填充内容
func (c *DemoConfig) UnmarshalByKey(key string, v any) error {
	cfg := c.Sub(key)
	if cfg != nil {
		return cfg.Unmarshal(v)
	}
	return nil
}

// ValueByKey 根据key获取内容，返回 any，找不到则返回 nil
func (c *DemoConfig) ValueByKey(key string) any {
	return c.Get(key)
}

// StrValueByKey 根据key获取字符串内容，找不到返回空字符串
func (c *DemoConfig) StrValueByKey(key string) string {
	return c.GetString(key)
}

// SetByKey 覆盖配置
func (c *DemoConfig) SetByKey(key string, v any) error {
	c.Set(key, v)
	return nil
}
