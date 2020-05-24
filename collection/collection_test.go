package collection

import (
	"fmt"
	"testing"
)

func ExampleSliceRemoveDuplicate() {
	//package main
	//
	//import (
	//	"fmt"
	//"github.com/ourlang/noUtil/collection"
	//)
	//func main()	{
	//	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	//	newSlice := SliceRemoveDuplicate(arr)
	//	fmt.Println(newSlice) //[hello hi world china]
	//}
}

func TestSliceRemoveDuplicate(t *testing.T) {
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	newSlice := SliceRemoveDuplicate(arr)
	fmt.Println(newSlice) //[hello hi world china]

	var arr1 = []int{1, 2, 4, 2, 5, 6, 8, 5, 9, 23}
	newSlice = SliceRemoveDuplicate(arr1)
	fmt.Println(newSlice) //[1 2 4 5 6 8 9 23]

	var arr2 = []float64{1.12, 2.45, 3.14, 5.43, 76.89, 3.14}
	newSlice = SliceRemoveDuplicate(arr2)
	fmt.Println(newSlice) //[1.12 2.45 3.14 5.43 76.89]
}
