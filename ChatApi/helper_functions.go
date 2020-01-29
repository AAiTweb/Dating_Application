package ChatApi

import (
	"fmt"
	"time"
)

type Month int

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func MonthConverter(month int) string {
	switch month {
	case int(January):
		return "January"
	case int(February):
		return "February"
	case int(March):
		return "March"
	case int(April):
		return "April"
	case int(May):
		return "May"
	case int(June):
		return "June"
	case int(July):
		return "July"
	case int(August):
		return "August"
	case int(September):
		return "September"
	case int(October):
		return "October"
	case int(November):
		return "November"
	case int(December):
		return "December"

	}
	return ""
}
func ClockFormatter(hour, min int) string {
	Meridien := ""
	if hour < 12 {
		Meridien = "AM"
	} else {
		Meridien = "PM"
	}
	hour = hour % 12
	return fmt.Sprintf("%d:%d %s", hour, min, Meridien)
}
func MessageSendTimeChanger(sendTime time.Time) string {
	t := ""
	rightNowTime := time.Now()
	yearDifference := sendTime.Year() - rightNowTime.Year()
	month := MonthConverter(int(sendTime.Month()))
	day := sendTime.Day()
	year := sendTime.Year()
	hour, min, _ := sendTime.Clock()
	formattedClock := ClockFormatter(hour, min)
	if yearDifference > 0 {
		t = fmt.Sprintf("%s %d %d, %s", month, day, year, formattedClock)
	} else {
		t = fmt.Sprintf("%s %d, %s", month, day, formattedClock)
	}
	return t
}
