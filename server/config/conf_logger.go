package config

type Logger struct {
	Level       string `yaml:"level"`
	Perfix      string `yaml:"perfix"`
	Director    string `yaml:"director"`
	ShowLine    bool   `yaml:"show_line"` //是否显示行号
	LogInConsol bool   `yaml:"log-in-consol"`
}
