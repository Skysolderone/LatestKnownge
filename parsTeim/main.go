package main

import (
	"fmt"
	"time"
)

func main1() {
	// Set the specific time for today
	// layout := "2006-01-02 15:04:05"
	// timeString := fmt.Sprintf("%s 09:05:25", time.Now().Format("2006-01-02"))

	// t, err := time.Parse(layout, timeString)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Convert the time to a Unix timestamp
	// timestamp := t.Unix()

	// fmt.Println("The Unix timestamp is:", timestamp)

	// now := time.Now()

	// // Get the hours, minutes, and seconds
	// hour, min, sec := now.Clock()

	// // Create a time object for today's date with the current time
	// t := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, time.UTC)

	// // Get the Unix timestamp
	// timestamp := t.Unix()

	// fmt.Println("The Unix timestamp is:", timestamp)

	// Convert the timestamp to a time.Time object
	timess := int64(1729671172)
	t := time.Unix(timess, 0)

	// Get the hours, minutes, and seconds
	hour, min, sec := t.Clock()

	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", hour, min, sec)
}
