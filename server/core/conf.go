package core

import (
	"Blog/config"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal %v", err)
	}
	log.Println("config yamlFile load Init success")
	fmt.Println(c)

}
