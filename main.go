package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: datacalc <+N> or <-N>        <---provide  one number with '-' or '+' as a prefix")
		return
	}

	//determine the sign
	arg := os.Args[1]
	var sign int

	var num string
	switch {
	case strings.HasPrefix(arg, "+"):
		sign = 1

		//slicing everything after the prefix
		num = arg[1:]
	case strings.HasPrefix(arg, "-"):
		sign = -1

		//slicing everything after the prefix
		num = arg[1:]
	default:
		fmt.Printf("Invalid argument: %s   <---provide  one number with '-' or '+' as a prefix\n", arg)
		return
	}

	// Convert the argument to an integer
	N, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid number:", num)
		return
	}

	// Apply the sign to N
	N = N * sign

	// Get the current date and time
	currentTime := time.Now()
	// Define the desired output format
	const layout_sec = "2006-01-02 15:04 seconds: 05"
	const layout_min = "2006-01-02 15:04"
	const layout_day = "2006-01-02"

	// Calculate the future dates
	futureSeconds := currentTime.Add(time.Duration(N) * time.Second)
	futureMinutes := currentTime.Add(time.Duration(N) * time.Minute)
	futureHours := currentTime.Add(time.Duration(N) * time.Hour)
	futureDays := currentTime.AddDate(0, 0, N)
	futureMonths := currentTime.AddDate(0, N, 0)
	futureYears := currentTime.AddDate(N, 0, 0)

	// Print the results with layout_sec
	fmt.Printf("\nCurrent date and time: %s\n\n", currentTime.Format(layout_sec))
	fmt.Printf(" %+d seconds: %s\n", N, futureSeconds.Format(layout_sec))
	fmt.Printf(" %+d minutes: %s\n", N, futureMinutes.Format(layout_sec))
	// Print the results with layout_min
	fmt.Printf(" %+d hours: %s\n", N, futureHours.Format(layout_min))
	fmt.Printf(" %+d days: %s\n", N, futureDays.Format(layout_min))
	//// Print the results with layout_day
	fmt.Printf(" %+d months: %s\n", N, futureMonths.Format(layout_day))
	fmt.Printf(" %+d years: %s\n\n", N, futureYears.Format(layout_day))
}
