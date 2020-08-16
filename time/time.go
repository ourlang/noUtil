// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package time provide common time and date operation common method.
package timeutil

import "time"

//Generating cron expressions by date time
func GetCronStr(date time.Time) string {
	dateFormat := "05 04 15 02 01 ?"
	format := date.Format(dateFormat)
	return format
}
