package conf

// Config configure args
type Config struct {
	Host           string `json:"Host"`
	Port           int    `json:"Port"`
	Incluster      string `json:"incluster"`
	Namespace      string `json:"Namespace"`
	IstioNamespace string `json:"IstioNamespace"`
}

var _cfg *Config

func SetConfig(cfg *Config) {
	_cfg = cfg
}

func GetConfig() *Config {
	return _cfg
}
