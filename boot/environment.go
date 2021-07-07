package boot

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

//应用环境变量
type Environment struct {
	//服务配置
	Server struct {
		//端口号
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	}
}

var env Environment

//初始化,分别从 ./application.yml ./config/application.yml 读取配置文件
func loadEnvironment() {

	content, err := ioutil.ReadFile("./application.yml")

	//TODO 如果没有，那么要有一个默认的配置
	if err != nil {
		log.Fatalf("读取application.yml错误: %v", err)
	}

	if yaml.Unmarshal(content, &env) != nil {
		log.Fatalf("解析application.yml出错: %v", err)
	}

	jsonInfo, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		log.Fatalf("解析application.yml出错: %v", err)
	}

	fmt.Println("---系统配置信息---")
	fmt.Println(string(jsonInfo))
	fmt.Println("---------------")
}

//获取当前环境
func getEnvironment() Environment {
	return env //返回的是一个值对象，不能返回引用(指针)。防止对象被修改
}
