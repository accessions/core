package times

import (
	"fmt"
	"time"
)

func test()  {
	start := time.Now()
	format := time.Now().Format(time.RFC3339)
	println(format)

	time.Sleep(1e9) // 1秒
	time.Sleep(3e9) // 3秒
	time.Sleep(1e6) // 1毫秒
	time.Sleep(3e6) // 3毫秒

	sub := time.Now().Sub(start)
	fmt.Println("use time : ", sub)
}

package lib

import "time"

var (
	Week = map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}
)

func week() (weekDate time.Time) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

func CurrentWeek() (startWeekDate, endWeekDate time.Time) {
	return week(), week().AddDate(0, 0, 6)
}

func PrevWeek() (prevStartWeekDate, prevEndWeekDate time.Time) {
	return week().AddDate(0, 0, -7), week().AddDate(0, 0, -1)
}
