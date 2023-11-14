package echoext

type Options struct {
	Name     string `json:"name,omitempty"`
	Version  string `json:"version,omitempty"`
	BasePath string `json:"base_path,omitempty" yaml:"base-path,omitempty"`
	Port     int    `json:"port,omitempty"`
	TLSKey   string `json:"tls_key,omitempty" yaml:"tls-key,omitempty"`
	TLSPem   string `json:"tls_pem,omitempty" yaml:"tls-pem,omitempty"`

	authorMap map[string]Author
}

func (o *Options) SetName(name string) *Options {
	o.Name = name
	return o
}

func (o *Options) SetVersion(version string) *Options {
	o.Version = version
	return o
}

func (o *Options) SetAuthor(key string, author Author) *Options {
	if len(o.authorMap) == 0 {
		o.authorMap = make(map[string]Author)
	}
	o.authorMap[key] = author
	return o
}

func (o *Options) SetAuthorMap(authorMap map[string]Author) *Options {
	o.authorMap = authorMap
	return o
}

func (o *Options) GetAuthor(key string) Author {
	return o.authorMap[key]
}
