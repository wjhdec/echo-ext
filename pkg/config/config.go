package config

type Config interface {
	// Reload 重新加载配置
	Reload() error
	// UnmarshalByKey 根据key填充内容
	UnmarshalByKey(key string, v any) error
	// ValueByKey 根据key获取内容，返回 any，找不到则返回 nil
	ValueByKey(key string) any
	// StrValueByKey 根据key获取字符串内容，找不到返回空字符串
	StrValueByKey(key string) string
	// ConfigFileUsed 返回使用的配置文件路径
	ConfigFileUsed() string
	// SetByKey 覆盖配置
	SetByKey(key string, v any) error
}
