// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package time provide common time and date operation common method.
package timeutil

import (
	"strings"
	"time"
)

const (
	DEFAULTTIMEFORMAT = "2006-01-02 15:04:05"
	TIMEFORMAT        = "20060102150405"
)

// get current time
func GetNowTime() time.Time {
	return time.Now()
}

// Gets the formatted string of the time
func TimeFormat(time *time.Time, format string) string {
	var datePatterns = []string{
		// year
		"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
		"y", "06", //A two digit representation of a year   Examples: 99 or 03

		// month
		"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
		"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
		"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
		"F", "January", // A full textual representation of a month, such as January or March   January through December

		// day
		"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
		"j", "2", // Day of the month without leading zeros 1 to 31

		// week
		"D", "Mon", // A textual representation of a day, three letters Mon through Sun
		"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

		// time
		"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
		"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
		"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
		"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

		"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
		"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

		"i", "04", // Minutes with leading zeros    00 to 59
		"s", "05", // Seconds, with leading zeros   00 through 59

		// time zone
		"T", "MST",
		"P", "-07:00",
		"O", "-0700",

		// RFC 2822
		"r", "Mon, 02 Jan 2006 15:04:05 -0700",
	}
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)
	return time.Format(format)
}

func StringFormatTime(timeLayout string) int64 {
	theTime, _ := time.Parse(DEFAULTTIMEFORMAT, timeLayout)
	timeUnix := theTime.Unix()
	return timeUnix
}

//Generating cron expressions by date time
func GetCronStr(date time.Time) string {
	dateFormat := "05 04 15 02 01 ?"
	format := date.Format(dateFormat)
	return format
}

//Get time string
func GetTimeString(t time.Time, timeLayout string) string {
	return t.Format(timeLayout)
}

// Timestamp to Seconds
func GetTimeUnix(t time.Time) int64 {
	return t.Unix()
}

// Timestamp to milliseconds
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// Time stamp to time
func GetTimeByInt(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

// Compare the two time sizes
func CompareTime(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// How many hours is the difference between the two
func GetHourDiffer(startTime, endTime string) float32 {
	var hour float32
	t1, err := time.ParseInLocation(DEFAULTTIMEFORMAT, startTime, time.Local)
	t2, err := time.ParseInLocation(TIMEFORMAT, endTime, time.Local)
	if err == nil && CompareTime(t1, t2) {
		diff := GetTimeUnix(t2) - GetTimeUnix(t1)
		hour = float32(diff) / 3600
		return hour
	}
	return hour
}

// Judge whether the current time is on the hour
func CheckHours() bool {
	_, m, s := GetNowTime().Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}
