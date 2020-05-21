// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// package collection is Golang Common methods of slice and array
// Including general operations such as slice and array de duplication, delete, and empty set
// For sorting and adding elements of slices and arrays, please use the official sort package related methods
package collection

import (
	"fmt"
	"reflect"
)

//Remove duplicate elements from collection
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
