package config

import (
	"log/slog"
)

var defaultConfig Config

func SetDefaultConfig(cfg Config) {
	defaultConfig = cfg
}

// checkNotNil 检查是否设置配置内容
func checkNotNil() {
	if defaultConfig == nil {
		panic("should init global config")
	}
}

func DefaultConfig() Config {
	return defaultConfig
}

func Reload() {
	checkNotNil()
	if err := defaultConfig.Reload(); err != nil {
		slog.Error("reload config error")
	}
}

func UnmarshalByKey(key string, v any) error {
	checkNotNil()
	return defaultConfig.UnmarshalByKey(key, v)
}

func ValueByKey(key string) any {
	checkNotNil()
	return defaultConfig.ValueByKey(key)
}

func StrValueByKey(key string) string {
	checkNotNil()
	return defaultConfig.StrValueByKey(key)
}

func ConfigFileUsed() string {
	checkNotNil()
	return defaultConfig.ConfigFileUsed()
}

func SetByKey(key string, v any) error {
	checkNotNil()
	return defaultConfig.SetByKey(key, v)
}
