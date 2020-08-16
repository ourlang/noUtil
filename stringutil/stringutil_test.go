package stringutil

import (
	"fmt"
	"testing"
)

func TestStringBuilder_AppendStrings(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	builder := NewStringBuilder()
	builder.AppendStrings(s1, s2, "china")
	//helloworldchina
	fmt.Println(builder.ToString())
}

func TestStringBuilder_Clear(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	builder := NewStringBuilder()
	builder.Append(s1).Append(s2).Append("中国")
	fmt.Println("builder clear before ", builder.ToString())
	builder.Clear()
	fmt.Println("builder clear after ", builder.ToString())
}

func TestSubString(t *testing.T) {
	str := "helloWorld"
	newStr := SubString(str, 2)
	//lloWorld
	fmt.Println(newStr)
}

func TestString(t *testing.T) {
	builder := NewStringBuilderString("hello")
	builder.AppendStrings("world", "good", "ok")
	//helloworldgoodok
	fmt.Println(builder.ToString())
	newMap := make(map[string]string, 5)
	newMap["hello"] = "world"
	newMap["good"] = "bi"
	newMap["our"] = "lang"
	//{"good":"bi","hello":"world","our":"lang"}
	fmt.Println(String(newMap))
}
