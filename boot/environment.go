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
	//日志配置
	Logging LoggingConfigurer
	//服务配置
	Server Server
}

//服务配置
type Server struct {
	//端口号
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

//日志配置
type LoggingConfigurer struct {
	Level       string `yaml:"level"`         //warn,info,error,debug
	FilePath    string `yaml:"file-path"`     // ./logs/app.log
	FileMaxSize int    `yaml:"file-max-size"` // 10 #10M
	MaxHistory  int    `yaml:"max-history"`   //30
	Compress    bool   `yaml:"compress"`
	LocalTime   bool   `yaml:"local-time"`
	MaxBackups  int    `yaml:"max-backups"` //10
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

	fmt.Println(string(jsonInfo))
}

//获取当前环境
func Env() Environment {
	return env //返回的是一个值对象，不能返回引用(指针)。防止对象被修改
}
