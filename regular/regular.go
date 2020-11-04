// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package regular is Application of common regular expressions
package regular

import (
	"fmt"
	"regexp"
)

//email verify
func VerifyEmailFormat(email string) bool {
	//Matching email
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func CheckPasswordLevel(ps string) error {
	if len(ps) < 8 {
		return fmt.Errorf("password len is < 8")
	}
	num := `[0-9]{1}`
	az := `[A-Za-z]{1}`
	//symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(az, ps); !b || err != nil {
		return fmt.Errorf("password need A_Z :%v", err)
	}
	return nil
}

func CheckUserName(username string) (err error) {
	zz := `^\w{6,30}$`
	if b, err := regexp.MatchString(zz, username); !b || err != nil {
		return fmt.Errorf("illegal user name:%v", err)
	}
	zz = `^\d+$`
	if b, err := regexp.MatchString(zz, username); b || err != nil {
		return fmt.Errorf("illegal user name:%v", err)
	}
	return
}

func CheckNickName(nikeName string) (err error) {
	re := regexp.MustCompile("[\u0020-\u002F]|[\u003A-\u0040]|[\u005B-\u0060]|[\u00A0-\u00BF]")
	if err := re.MatchString(nikeName); err == true {
		return fmt.Errorf("illegal user nickname:%v", err)
	}
	return
}
