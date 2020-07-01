// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package time provide common time and date operation common method.
package time

/**
 * 通过输入指定日期时间生成cron表达式
 * @author 林敏
 * @param date 时间参数
 * @return cron表达式
 */

func GetCronStr(date time.Time) string {
	//定义时间格式化格式
	dateFormat := "05 04 15 02 01 ?"
	format := date.Format(dateFormat)
	return format
}
