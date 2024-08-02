package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site-info"`
	QQ       QQ       `yaml:"qq"`
	JWT      JWT      `yaml:"jwt"`
	Email    Email    `yaml:"email"`
	QiNiu    QiNiu    `yaml:"qiniu"`
	Upload   Upload   `yaml:"upload"`
}
