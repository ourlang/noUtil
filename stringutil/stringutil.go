// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package stringutil is common string tools
// A supplement to strings package
package stringutil

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
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

func RandomString(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
