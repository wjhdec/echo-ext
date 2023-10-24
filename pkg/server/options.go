package server

import (
	"io"
	"os"

	"github.com/wjhdec/echo-ext/pkg/config"
)

// ServerOptions 逻辑服务配置
type ServerOptions struct {
	Name      string
	Version   string
	logWriter io.Writer
	authMap   map[string]Author
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Name:      "",
		Version:   "",
		logWriter: os.Stdout,
		authMap:   make(map[string]Author),
	}
}

func (s *ServerOptions) SetName(name string) *ServerOptions {
	s.Name = name
	return s
}

func (s *ServerOptions) SetVersion(version string) *ServerOptions {
	s.Version = version
	return s
}

func (s *ServerOptions) AddAuth(key string, author Author) *ServerOptions {
	if len(s.authMap) == 0 {
		s.authMap = make(map[string]Author)
	}
	s.authMap[key] = author
	return s
}

func (s *ServerOptions) AddAuthMap(authMap map[string]Author) *ServerOptions {
	if len(s.authMap) == 0 {
		s.authMap = make(map[string]Author)
	}
	for k, a := range authMap {
		s.authMap[k] = a
	}
	return s
}

func (s *ServerOptions) Author(key string) Author {
	return s.authMap[key]
}

func NewConfigOptions(name string) *ConfigOptions {
	opt := new(ConfigOptions)
	err := config.UnmarshalByKey(name, opt)
	if err != nil {
		panic(err)
	}
	return opt
}

// ConfigOptions 文件配置
type ConfigOptions struct {
	BasePath string `json:"base_path"`
	Port     int    `json:"port"`
	TLSKey   string `json:"tls_key"`
	TLSPem   string `json:"tls_pem"`
}
