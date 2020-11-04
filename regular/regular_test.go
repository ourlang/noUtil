package regular

import (
	"fmt"
	"testing"
)

func ExampleVerifyEmailFormat() {
	format := VerifyEmailFormat("ourlang@foxmail.com")
	fmt.Println(format)
}

func TestVerifyEmailFormat(t *testing.T) {
	format := VerifyEmailFormat("ourlang@foxmail.com")
	fmt.Println(format)
}

func ExampleVerifyMobileFormat() {
	format := VerifyMobileFormat("15912338765")
	fmt.Println(format)
}

func ExampleCheckPasswordLevel() {
	err := CheckPasswordLevel("12345678")
	if err != nil {
		fmt.Println(err)
	}
	err = CheckPasswordLevel("hello!@43world(*")
	if err != nil {
		fmt.Println(err)
	}
}
