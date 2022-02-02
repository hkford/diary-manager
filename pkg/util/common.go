package util

var Days = []string{"Sat", "Sun", "Mon", "Tue", "Wed", "Thur", "Fri"}
var MonthNames = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
var DayLengths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func IsLeapYear(year int64) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}
