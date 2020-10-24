package date

import "time"

var apiDateLayout = "2006-01-02 15:04:05.000"

// GetNow func with UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString func with UTC
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetFutureTime func
func GetFutureTime(year int, month int, day int) string {
	return GetNow().AddDate(year, month, day).Format(apiDateLayout)
}

// ParseDate func
// parse string to date
func ParseDate(date string) time.Time {
	t,_ := time.Parse(apiDateLayout,date)
	return t 
}