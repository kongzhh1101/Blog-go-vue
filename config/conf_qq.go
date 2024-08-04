package config

type QQ struct {
	AppID    string `json:"app_id" yaml:"app-id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"`
}

func (c *QQ) GetPath() string {
	return "qq"
}
