package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"`
	Password         string `json:"passwrod" yaml:"password"`
	DefaultFromEmail string `json:"default-from-email" yaml:"default-from-email"`
	UserSSL          bool   `json:"user-ssl" yaml:"user-ssl"`
	UserTLS          bool   `json:"user-tls" yaml:"user-tls"`
}
