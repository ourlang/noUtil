package collection

import (
	"fmt"
	"testing"
)

//params: a: slice object，example:[]string, []int, []float64, ...

//return:  []interface{}: New slice object with duplicate elements removed
func TestSliceRemoveDuplicate(t *testing.T) {
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	newSlice := SliceRemoveDuplicate(arr)
	fmt.Println(newSlice) //[hello hi world china]
}
