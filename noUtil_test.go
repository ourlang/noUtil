package noUtil

import (
	"fmt"
	"testing"
)

func TestCommonTool(t *testing.T) {

	i, _ := Tools.StrToInt("123", 0)
	fmt.Println(i)
	dataBaseObj := DataBaseConnection{
		IpAddress:    "192.168.0.100",
		Port:         3306,
		UserName:     "root",
		Password:     123456,
		DataBaseName: "user_test",
	}
	ret, err := Tools.StructToMap(&dataBaseObj)
	if err != nil {
		fmt.Println("struct to map err:", err.Error())
	}
	for k, v := range ret {
		fmt.Println(k, v)
	}
}
