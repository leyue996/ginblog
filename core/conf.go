package core

import (
	"fmt"
	"ginblog/config"
	"ginblog/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Initconfig用来把yaml的数据传递给变量
func Initconfig() {
	const configFile = "settings.yaml"
	c := &config.Config{}
	yamlconf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("configFile打开错误:%s", err))
	}

	err = yaml.Unmarshal(yamlconf, c)
	if err != nil {
		log.Fatalf("ymal数据给c时错误:%s", err)
	}
	log.Println("config数据初始化成功")
	fmt.Println(c)
	global.Config = c
}
