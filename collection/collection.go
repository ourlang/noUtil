// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// package collection is Golang Common methods of slice and array
// Including general operations such as slice and array de duplication, delete, and empty set
// For sorting and adding elements of slices and arrays, please use the official sort package related methods
package collection

import (
	"errors"
	"fmt"
	"reflect"
)

//Remove duplicate elements from collection

//params: a: slice object，example:[]string, []int, []float64, ...
//return:  []interface{}: New slice object with duplicate elements removed
func SliceRemoveDuplicate(a interface{}) (ret []interface{}) {
	if reflect.TypeOf(a).Kind() != reflect.Slice {
		fmt.Printf("The parameter passed in is not a slice but %T\n", a)
		return ret
	}
	va := reflect.ValueOf(a)
	m := make(map[interface{}]bool)
	for i := 0; i < va.Len(); i++ {
		sliceVal := va.Index(i).Interface()
		if _, ok := m[sliceVal]; !ok {
			ret = append(ret, sliceVal)
			m[sliceVal] = true
		}
	}
	return ret
}

// InsertSlice insert a element to slice in the specified index
// Note that original slice will not be modified
func InsertSlice(slice interface{}, index int, value interface{}) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("target isn't a slice")
	}
	if index < 0 || index > v.Len() || reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return nil, errors.New("param is invalid")
	}

	dst := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)

	// add the element to the end of slice
	if index == v.Len() {
		dst = reflect.AppendSlice(dst, v.Slice(0, v.Len()))
		dst = reflect.Append(dst, reflect.ValueOf(value))
		return dst.Interface(), nil
	}

	dst = reflect.AppendSlice(dst, v.Slice(0, index+1))
	dst = reflect.AppendSlice(dst, v.Slice(index, v.Len()))
	dst.Index(index).Set(reflect.ValueOf(value))
	return dst.Interface(), nil
}

// DeleteSliceE deletes the specified subscript element from the slice
// Note that original slice will not be modified
func DeleteSliceE(slice interface{}, index int) (interface{}, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("target isn't a slice")
	}
	if v.Len() == 0 || index < 0 || index > v.Len()-1 {
		return nil, errors.New("param is invalid")
	}

	dst := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	dst = reflect.AppendSlice(dst, v.Slice(0, index))
	dst = reflect.AppendSlice(dst, v.Slice(index+1, v.Len()))
	return dst.Interface(), nil
}

// GetEleIndexesSliceE finds all indexes of the specified element in a slice
func GetEleIndexesSliceE(slice interface{}, value interface{}) ([]int, error) {
	// check params
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("target isn't a slice")
	}

	var indexes []int
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == value {
			indexes = append(indexes, i)
		}
	}
	return indexes, nil
}

//Returns true if this array、slice、map contains the specified element.
//More formally, returns true if and only if this list contains
//at least one element e such that

//@params:list---> slice、array、map，example:[]string, []int, []float64, map[string]string……
//@params:e--->An element in a slice array map.example:"hello",123……
func Contains(list interface{}, e interface{}) (bool, error) {
	listType := reflect.TypeOf(list).Kind()
	listValue := reflect.ValueOf(list)
	if listType == reflect.Slice || listType == reflect.Array || listType == reflect.Map {
		switch listType {
		case reflect.Slice, reflect.Array:
			for i := 0; i < listValue.Len(); i++ {
				if listValue.Index(i).Interface() == e {
					return true, nil
				}
			}
		case reflect.Map:
			if listValue.MapIndex(reflect.ValueOf(e)).IsValid() {
				return true, nil
			}
		}

		return false, errors.New(reflect.ValueOf(e).String() + " not in list")
	} else {
		return false, errors.New(listValue.String() + " It is not any type in array, slice or map ")
	}
}

//Gets the union of two slices
func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//Get the intersection of two slices
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//Get the difference set of two slices
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func StructToMap(obj interface{}) (result map[string]interface{}, err error) {
	k := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if k.Kind() != reflect.Ptr {
		err = fmt.Errorf("type must be a pointer")
		return
	}

	if k.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("element type must be a struct")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()
	result = make(map[string]interface{})
	for i := 0; i < k.Elem().NumField(); i++ {
		name := k.Elem().Field(i).Name
		field := v.Elem().Field(i)
		switch name {
		case "ID", "CreatedAt", "UpdatedAt", "DeletedAt":
			continue
		}
		switch field.Kind() {
		case reflect.Slice, reflect.Struct, reflect.Ptr:
			continue
		default:
			result[name] = field.Interface()
		}
	}
	return
}
