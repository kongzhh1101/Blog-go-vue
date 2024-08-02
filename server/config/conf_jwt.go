package config

type JWT struct {
	Secret string `json:"secret" yaml:"secret"`
	Expire int64  `json:"expire" yaml:"expire"`
	Issuer string `json:"issure" yaml:"issuer"`
}
