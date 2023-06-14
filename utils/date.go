package utils

import (
	"fmt"
	"time"
)

func DateFormat(dateData time.Time) (string) {

	dateFormat := dateData.Local()
	year := dateFormat.Year()
	month := dateFormat.Month().String()
	day := dateFormat.Day()
	time := dateFormat.Format("15:04")

	return fmt.Sprintf("%s %d, %d เมื่อ %s", month, day, year, time)
}