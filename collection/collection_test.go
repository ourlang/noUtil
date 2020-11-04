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

func TestContains(t *testing.T) {
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	flag, err := Contains(arr, "hello")
	fmt.Println(flag, err)

	map1 := make(map[string]string, 0)
	map1["hello"] = "golang"
	map1["world"] = "china"
	flag, err = Contains(map1, "hello1")
	fmt.Println(flag, err)
}

func TestGetEleIndexesSliceE(t *testing.T) {
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	indexSlice, err := GetEleIndexesSliceE(arr, "hi")
	fmt.Println(indexSlice, err)
}

func ExampleDifference() {
	slice1 := []string{"hello", "china", "good"}
	slice2 := []string{"hello", "china", "happy"}
	difSlice := Difference(slice1, slice2)
	fmt.Println(difSlice) //[good]

	unionSlice := Union(slice1, slice2)
	fmt.Println(unionSlice) //[hello china good happy]

	intersectionSlice := Intersect(slice1, slice2)
	fmt.Println(intersectionSlice) //[hello china]
}

func ExampleStructToMap() {
	type Student struct {
		Name    string
		Age     int
		Address string
		Email   string
	}
	stu := Student{
		Name:    "ourlang",
		Age:     30,
		Address: "chengdu",
		Email:   "ourlang@foxmail.com",
	}
	tempMap, err := StructToMap(&stu)
	if err != nil {
		panic(err.Error())
	}
	for s, v := range tempMap {
		fmt.Println(s, v)
	}

}
