package core

import (
	"Blog/config"
	"Blog/global"
	"fmt"
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const ConfigFile = "config.yaml"

// InitConf 读取yaml文件的配置
func InitConfig() {
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
	global.Config = c
}

// Updateconfig 更新yaml文件的配置
func UpdateConfig() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	log.Println("config Updateconfig success")
	return nil
}
