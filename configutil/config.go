// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// package configutil is Golang Configuration Solution
// Supports reading in JSON, toml, yaml, HCl and Java property configuration files
package configutil

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

//Read configuration file function (need to be implemented by reflection)
//
//configFilePath:配置文件路径
//
//configFileName:配置文件路径下的文件名
//
//configFileType:配置文件类型，如JSON, toml, yaml, HCl，
//
//obj:把指定文件的内容读取到对应的结构体
func ReadConfigFile(configFilePath, configFileName, configFileType string, obj interface{}) {
	config := viper.New()
	//配置文件名（不带扩展名）
	config.SetConfigName(configFileName)
	//在项目中查找配置文件的路径，可以使用相对路径，也可以使用绝对路径
	config.AddConfigPath(configFilePath)
	//多次调用以添加多个搜索路径
	//viper.AddConfigPath("D:/go_project/src/github.com/ourlang/demo/utils")
	//设置文件类型，这里是yaml文件
	config.SetConfigType(configFileType)
	//定义用于接收配置文件的变量
	//查找并读取配置文件
	err := config.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//声明一个读取配置文件的map
	m := make(map[string]interface{})
	if err := config.Unmarshal(&m); err != nil { // 读取配置文件转化成对应的结构体错误
		panic(fmt.Errorf("read config file to struct err: %s \n", err))
	}
	mapstructure.Decode(m, &obj)
}
