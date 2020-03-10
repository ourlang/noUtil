package configutil

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"testing"
)

func ExampleReadConfigFile() {
	//定义接收配置文件的结构体
	type DataBaseConnection struct {
		IpAddress    string
		Port         int
		UserName     string
		Password     int
		DataBaseName string
	}
	m := ReadConfigFile("D:/go_project/src/github.com/ourlang/demo/config", "appConfig", "yaml")
	var config DataBaseConnection
	mapstructure.Decode(m, &config)
	fmt.Println(config) //{127.0.0.1 3306 root 123456 go_test}
}

func TestReadConfigFile(t *testing.T) {
	ExampleReadConfigFile()
}
