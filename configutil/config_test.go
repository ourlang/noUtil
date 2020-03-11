package configutil

import (
	"fmt"
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
	var db DataBaseConnection
	ReadConfigFile("E:/go_project/SentaSystemManage/config/", "ApplicationConfig", "yaml", &db)
	fmt.Println(db) //{127.0.0.1 3306 root 123456 go_test}
}

func TestReadConfigFile(t *testing.T) {
	ExampleReadConfigFile()
}
