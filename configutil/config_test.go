package configutil

import "testing"

//定义接收配置文件的结构体
type DataBaseConnection struct {
	IpAddress    string
	Port         int
	UserName     string
	Password     int
	DataBaseName string
}

func ExampleReadConfigFile() {
	ReadConfigFile("D:/go_project/src/github.com/ourlang/demo/config", "appConfig", "yaml", DataBaseConnection{})
}

func TestReadConfigFile(t *testing.T) {
	ExampleReadConfigFile()
}
