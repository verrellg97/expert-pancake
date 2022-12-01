package util

import (
	"strconv"
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
	dateValue, _ = time.Parse("2006-01-02", value)
	return dateValue
}

func DateToString(value time.Time) string {
	return value.Format("2006-01-02")
}

func StringToBigInt(value string) int64 {
	strValue, _ := strconv.ParseInt(value, 10, 64)
	return strValue
}

func BigIntToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func BoolToString(value bool) string {
	if value {
		return "1"
	} else {
		return "0"
	}
}
