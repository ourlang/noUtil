// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package stringutil is common string tools
// A supplement to strings package
package stringutil

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/*
  StringBuilder struct.
	  Usage:
		builder := NewStringBuilderString("hello")
		builder.AppendStrings("World", "Study", "Golang")
		fmt.Println(builder.ToString())
	  print:
		helloWorldStudyGolang
*/
type StringBuilder struct {
	builderStruct strings.Builder
}

func NewStringBuilder() *StringBuilder {
	var builder StringBuilder
	return &builder
}

func NewStringBuilderString(str string) *StringBuilder {
	var builder StringBuilder
	builder.builderStruct.WriteString(str)
	return &builder
}

//Splice a string into an existing StringBuilder
func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.builderStruct.WriteString(s)
	return builder
}

//Splicing multiple strings into an existing StringBuilder
func (builder *StringBuilder) AppendStrings(ss ...string) *StringBuilder {
	for i := range ss {
		builder.builderStruct.WriteString(ss[i])
	}
	return builder
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func (builder *StringBuilder) Replace(old, new string, n int) *StringBuilder {
	buildStr := builder.ToString()
	str := strings.Replace(buildStr, old, new, n)
	builder.Clear()
	builder.builderStruct.WriteString(str)
	return builder
}

func (builder *StringBuilder) Clear() *StringBuilder {
	var newBuilder strings.Builder
	builder.builderStruct = newBuilder
	return builder
}

// String returns the accumulated string
func (builder *StringBuilder) ToString() string {
	return builder.builderStruct.String()
}

// Substr returns a string of length length from the start position
func SubString(s string, start int, strLength ...int) string {
	charList := []rune(s)
	l := len(charList)
	length := 0
	end := 0

	if len(strLength) == 0 {
		length = l
	} else {
		length = strLength[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}

	return string(charList[start:end])
}

// String returns a string of any type
func String(obj interface{}) string {
	switch val := obj.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr, reflect.Struct, reflect.Map:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}
	return fmt.Sprintf("%v", obj)
}

//Generate random string
func GetRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()+[]{}/<>;:=.,?"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

//Generate random captcha
func GetRandomInt(n int) string {
	const letterBytes = "0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

//Amount in words
func AmountToCN(pMoney float64, pRound bool) string {
	var numberUpper = []string{"壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖", "零"}
	var unit = []string{"分", "角", "圆", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟"}
	var regex = [][]string{
		{"零拾", "零"}, {"零佰", "零"}, {"零仟", "零"}, {"零零零", "零"}, {"零零", "零"},
		{"零角零分", "整"}, {"零分", "整"}, {"零角", "零"}, {"零亿零万零元", "亿元"},
		{"亿零万零元", "亿元"}, {"零亿零万", "亿"}, {"零万零元", "万元"}, {"万零元", "万元"},
		{"零亿", "亿"}, {"零万", "万"}, {"拾零圆", "拾元"}, {"零圆", "元"}, {"零零", "零"}}
	str, digitUpper, unitLen, round := "", "", 0, 0
	if pMoney == 0 {
		return "零"
	}
	if pMoney < 0 {
		str = "负"
		pMoney = math.Abs(pMoney)
	}
	if pRound {
		round = 2
	} else {
		round = 1
	}
	digitByte := []byte(strconv.FormatFloat(pMoney, 'f', round+1, 64)) //注意币种四舍五入
	unitLen = len(digitByte) - round

	for _, v := range digitByte {
		if unitLen >= 1 && v != 46 {
			s, _ := strconv.ParseInt(string(v), 10, 0)
			if s != 0 {
				digitUpper = numberUpper[s-1]

			} else {
				digitUpper = "零"
			}
			str = str + digitUpper + unit[unitLen-1]
			unitLen = unitLen - 1
		}
	}
	for i := range regex {
		reg := regexp.MustCompile(regex[i][0])
		str = reg.ReplaceAllString(str, regex[i][1])
	}
	if string(str[0:3]) == "元" {
		str = string(str[3:])
	}
	if string(str[0:3]) == "零" {
		str = string(str[3:])
	}
	return str
}

func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := SubString(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			result = SubString(str, 0, 3) + "****" + SubString(str, 7, 11)
		} else {

			nameRune := []rune(str)
			lens := len(nameRune)
			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}
