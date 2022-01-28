// must not import a package called "init", so package name is a little bit long.
// https://go-review.googlesource.com/c/go/+/287494/
package initialize

import (
	"math"
)

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
	days := []string{"Sat", "Sun", "Mon", "Tue", "Wed", "Thur", "Fri"}
	second_term := int64(math.Floor(float64(26 * (m + 1) / 10)))
	Y := y % 100
	forth_term := int64(math.Floor(float64(Y / 4)))
	C := int64(math.Floor(float64(y / 100)))
	Gamma := 5*C + int64(math.Floor(float64(C/4)))
	h := (d + second_term + Y + forth_term + Gamma) % 7
	return days[h]
}
