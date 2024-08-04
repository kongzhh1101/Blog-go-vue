package config

type SiteInfo struct {
	Title   string `yaml:"title" json:"title"`
	Version string `yaml:"version" json:"version"`
	Email   string `yaml:"email" json:"email"`
}
