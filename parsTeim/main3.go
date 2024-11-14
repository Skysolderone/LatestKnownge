package main

import (
	"fmt"
	"time"
)

func getCountDown(ts []uint, t int64) int64 {
	tu := time.Unix(t, 0)

	nowOfWeek := uint(time.Now().Weekday())
	if nowOfWeek == 0 {
		nowOfWeek = 7
	}
	// fmt.Println(nowOfWeek)
	var s uint = 0
	var next uint = 0
	var date uint = 0
	for _, v := range ts {
		if v == nowOfWeek {
			date = v
		}

		if (nowOfWeek - v) > s {

			s = nowOfWeek - v
			next = v
		}
	}
	if next == 0 {
		next = nowOfWeek
	}

	if date != 0 && date == nowOfWeek {
		fmt.Println(0)
		// 判断是否已经开始
		now := time.Now()
		hour, min, sec := now.Clock()
		hour2, min2, sec2 := tu.Clock()
		// fmt.Println(hour2)
		fmt.Println(now.Format("2006-01-02"))
		fmt.Println(tu.Format("2006-01-02"))
		switch hour == hour2 {
		case true:
			switch min == min2 {
			case true:
				if sec < sec2 && now.Format("2006-01-02") == tu.Format("2006-01-02") {
					// nextDay := time.Date(.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())

					// DayTimestamp := nextDay.Unix()
					return t
				} else if hour < hour2 {
					nextDay := time.Date(now.Year(), now.Month(), now.Day(), hour2, min2, sec2, 0, now.Location())

					DayTimestamp := nextDay.Unix()
					return DayTimestamp
				}
			case false:
				if min < min2 && now.Format("2006-01-02") == tu.Format("2006-01-02") {
					// nextDay := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())

					// DayTimestamp := nextDay.Unix()
					return t
				} else if hour < hour2 {
					nextDay := time.Date(now.Year(), now.Month(), now.Day(), hour2, min2, sec2, 0, now.Location())

					DayTimestamp := nextDay.Unix()
					return DayTimestamp
				}
			}
		case false:
			if hour < hour2 && now.Format("2006-01-02") == tu.Format("2006-01-02") {
				// nextDay := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, now.Location())

				// DayTimestamp := nextDay.Unix()
				return t
			} else if hour < hour2 {
				nextDay := time.Date(now.Year(), now.Month(), now.Day(), hour2, min2, sec2, 0, now.Location())

				DayTimestamp := nextDay.Unix()
				return DayTimestamp
			}

		}

	}

	if next < nowOfWeek {
		fmt.Println(1)
		// 说明本周已经过去 为下周
		hour, min, sec := tu.Clock()
		now := time.Now()
		day := 7 - nowOfWeek
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+(int(day)+int(next)), hour, min, sec, 0, now.Location())

		DayTimestamp := nextDay.Unix()
		return DayTimestamp
	} else if next == nowOfWeek {
		fmt.Println(3)
		hour, min, sec := tu.Clock()
		now := time.Now()
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+7, hour, min, sec, 0, now.Location())
		// fmt.Println("000000")
		DayTimestamp := nextDay.Unix()
		return DayTimestamp
	} else {
		fmt.Println(2)
		hour, min, sec := tu.Clock()
		now := time.Now()
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+int(next-nowOfWeek), hour, min, sec, 0, now.Location())
		// fmt.Println("000000")
		DayTimestamp := nextDay.Unix()
		return DayTimestamp
	}
}

func main() {
	// 1730427540
	t := int64(1729649340)
	ts := []uint{3, 4}
	fmt.Println(getCountDown(ts, t))
}
