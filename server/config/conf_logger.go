package config

type Logger struct {
	Level       string `yaml:"level"`
	Perfix      string `yaml:"perfix"`
	Director    bool   `yaml:"director"`
	ShowLine    bool   `yaml:"show_line"` //是否显示行号
	LogInConsol bool   `yaml:"log_in_consol"`
}
