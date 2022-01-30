// must not import a package called "init", so package name is a little bit long.
// https://go-review.googlesource.com/c/go/+/287494/
package initialize

import (
	"fmt"
	"math"
)

var days = []string{"Sat", "Sun", "Mon", "Tue", "Wed", "Thur", "Fri"}
var monthNames = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
var dayLengths = []int64{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

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

// Get Gregorian calendar date using Zeller's congruence
// https://en.wikipedia.org/wiki/Zeller%27s_congruence
func GetDate(y int64, m int64, d int64) string {
	// January and February are counted as months 13 and 14 of the previous year
	if m <= 2 {
		m += 12
		y -= 1
	}

	second_term := int64(math.Floor(float64(26 * (m + 1) / 10)))
	Y := y % 100
	forth_term := int64(math.Floor(float64(Y / 4)))
	C := int64(math.Floor(float64(y / 100)))
	Gamma := 5*C + int64(math.Floor(float64(C/4)))
	h := (d + second_term + Y + forth_term + Gamma) % 7
	return days[h]
}

func GenerateDayFormat(y int64, m int64, d int64) string {
	f := fmt.Sprintf("%v,%v,%02v,%v\n\n", y, monthNames[m-1], d, GetDate(y, m, d))
	return f
}
