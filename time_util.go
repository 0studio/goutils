package goutils

import (
	"time"
)

var monthDays = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// 获取指定年月一共有多少天
func GetMonthDayCount(year, month int) int { // month[1~12]
	if month == 2 && IsLeapYear(year) {
		return 29
	}
	return monthDays[month]
}
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// 返回本月1号， 属于今年第几天
func GetYearDayOfCurrentMonthFirstDay(now time.Time) int32 {
	monthTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	now.Month()
	return int32(monthTime.YearDay())
}

// 返回本月最后一天， 属于今年第几天
func GetYearDayOfCurrentMonthLastDay(now time.Time) int32 {
	montDayCount := GetMonthDayCount(now.Year(), int(now.Month()))
	monthEndTime := time.Date(now.Year(), now.Month(), montDayCount, 0, 0, 0, 0, now.Location())
	return int32(monthEndTime.YearDay())
}
