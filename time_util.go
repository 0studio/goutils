package goutils

import (
	"strconv"
	"strings"
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
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// 返回本月1号， 属于今年第几天
func GetYearDayOfCurrentMonthFirstDay(now time.Time) int32 {
	monthTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return int32(monthTime.YearDay())
}

// 返回本月最后一天， 属于今年第几天
func GetYearDayOfCurrentMonthLastDay(now time.Time) int32 {
	montDayCount := GetMonthDayCount(now.Year(), int(now.Month()))
	monthEndTime := time.Date(now.Year(), now.Month(), montDayCount, 0, 0, 0, 0, now.Location())
	return int32(monthEndTime.YearDay())
}

/*
* TIME UTIL
 */

func IsSameDay(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return (y1 == y2 && m1 == m2 && d1 == d2)
}

func IsSameWeek(t1 time.Time, t2 time.Time) bool {
	year1, week1 := t1.ISOWeek()
	year2, week2 := t1.ISOWeek()
	return (year1 == year2 && week1 == week2)
}

func IsSameMonth(t1 time.Time, t2 time.Time) bool {
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()
	return (y1 == y2 && m1 == m2)
}

// 1=monday  7 =sunday
func GetWeekDay(t time.Time) int32 {
	if t.Weekday() == time.Sunday {
		return 7
	}
	return int32(t.Weekday())
}

func GetFormatDate(time time.Time) (timeStr string) {
	year := strconv.Itoa(time.Year())
	month := time.Month()
	day := time.Day()
	monthStr := ""
	dayStr := ""
	if month < 10 {
		monthStr = "0" + strconv.Itoa(int(month))
	} else {
		monthStr = strconv.Itoa(int(month))
	}
	if day < 10 {
		dayStr = "0" + strconv.Itoa(day)
	} else {
		dayStr = strconv.Itoa(day)
	}
	strs := []string{year, monthStr, dayStr}
	timeStr = strings.Join(strs, "-")
	return
}

// do some thing at 4:00 am.
// 获取从现在到 某一时间点 需要多少秒
func GetDurToNextDeadlineTime(now time.Time, hour int, min int, seconds int) (dur int) {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	if deadlineTime.Before(now) {
		tomorrow := now.Add(time.Hour * 24)
		deadlineTime = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), hour, min, seconds, 0, now.Location())
	}
	dur = int(deadlineTime.Sub(now).Seconds())
	return
}
func GetDurToNextDeadlineTimeDuration(now time.Time, hour int, min int, seconds int) (dur time.Duration) {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	if deadlineTime.Before(now) {
		tomorrow := now.Add(time.Hour * 24)
		deadlineTime = time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), hour, min, seconds, 0, now.Location())
	}
	dur = deadlineTime.Sub(now)
	return
}
func IsBeforeHMS(now time.Time, hour int, min int, seconds int) bool {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	return now.Before(deadlineTime)

}
func IsAfterHMS(now time.Time, hour int, min int, seconds int) bool {
	deadlineTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, seconds, 0, now.Location())
	return now.After(deadlineTime)

}
