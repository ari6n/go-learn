package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current date/time: ", time.Now())
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// fmt.Println("Saturday: ", time.Saturday)
	// fmt.Println("Yesterday's week-day: ", today+6)
	// fmt.Println("Current week-day: ", today)
	// fmt.Println("Tomorrow's week-day: ", today+1)
	switch time.Saturday {
	case today + 6:
		fmt.Println("Was yesterday.")
		fmt.Println("Test for several lines of code in switch's case")
		// fallthrough
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
