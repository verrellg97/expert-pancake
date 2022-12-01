package util

import (
	"time"
)

func WildCardString(keyword string) string {
	if keyword != "" {
		return "%" + keyword + "%"
	} else {
		return "%"
	}
}

func StringToDate(value string) time.Time {
	var dateValue time.Time
	dateValue, _ = time.Parse(DateLayoutYMD, value)
	return dateValue
}
