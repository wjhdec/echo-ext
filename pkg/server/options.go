package server

import "github.com/wjhdec/echo-ext/pkg/config"

func NewOptions(cfg config.Config, name string) *Options {
	opt := new(Options)
	err := cfg.UnmarshalByKey(name, opt)
	if err != nil {
		panic(err)
	}
	return opt
}

type Options struct {
	BasePath string `json:"base_path"`
	Port     int    `json:"port"`
	TLSKey   string `json:"tls_key"`
	TLSPem   string `json:"tls_pem"`
}
