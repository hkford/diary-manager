// must not import a package called "init", so package name is a little bit long.
// https://go-review.googlesource.com/c/go/+/287494/
package initialize

import (
	"fmt"
	"math"
	"mydiary/pkg/workspace"
	"strconv"
)

var days = []string{"Sat", "Sun", "Mon", "Tue", "Wed", "Thur", "Fri"}
var monthNames = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
var dayLengths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

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

func WriteMonthTemplate(ws workspace.Workspace, month int64) {
	var template = make([]byte, 0, 700)
	year, err := strconv.ParseInt(ws.DiaryDir, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse DiaryDir into int64")
	}
	days := dayLengths[month-1]
	if ws.IsLeap {
		days += 1
	}
	for d := 1; d <= days; d++ {
		dayFormat := GenerateDayFormat(year, month, int64(d))
		template = append(template, []byte(dayFormat)...)
	}
	filename := fmt.Sprintf("%v/%v%02v.txt", ws.DiaryDir, ws.DiaryDir, strconv.FormatInt(month, 10))
	ws.Fs.WriteFile(filename, template, 0755)
}

func WriteYearTemplates(ws workspace.Workspace) {
	for m := 1; m <= 12; m++ {
		WriteMonthTemplate(ws, int64(m))
	}
}
