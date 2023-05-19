// must not import a package called "init", so package name is a little bit long.
// https://go-review.googlesource.com/c/go/+/287494/
package initialize

import (
	"fmt"
	"math"
	"mydiary/pkg/util"
	"mydiary/pkg/workspace"
	"strconv"
	"time"
)

// Get Gregorian calendar date using Zeller's congruence
// https://en.wikipedia.org/wiki/Zeller%27s_congruence
func getDate(y int64, m int64, d int64) string {
	// January and February are counted as months 13 and 14 of the previous year
	if m <= 2 {
		m += 12
		y -= 1
	}

	second_term := int64(math.Floor((26 * float64(m+1) / 10)))
	Y := y % 100
	forth_term := int64(math.Floor(float64(Y) / 4))
	C := int64(math.Floor(float64(y) / 100))
	Gamma := 5*C + int64(math.Floor(float64(C)/4))
	h := (d + second_term + Y + forth_term + Gamma) % 7
	return util.Days[h]
}

func GenerateDayFormat(y int64, m time.Month, d int64) string {
	f := fmt.Sprintf("%v,%v,%02v,%v\n\n", y, m, d, getDate(y, int64(m), d))
	return f
}

func WriteMonthTemplate(ws workspace.Workspace, month time.Month) error {
	var template = make([]byte, 0, 700)
	year, err := strconv.ParseInt(ws.DiaryDir, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse DiaryDir into int64")
		return err
	}
	days := util.DayLengths[month-1]
	if ws.IsLeap && month == 2 {
		days += 1
	}
	for d := 1; d <= days; d++ {
		dayFormat := GenerateDayFormat(year, month, int64(d))
		template = append(template, []byte(dayFormat)...)
	}
	filename := fmt.Sprintf("diaries/%v/%v%02v.txt", ws.DiaryDir, ws.DiaryDir, strconv.FormatInt(int64(month), 10))
	err = ws.Fs.WriteFile(filename, template, 0755)
	return err
}

func WriteYearTemplates(ws workspace.Workspace) error {
	for m := time.January; m <= time.December; m++ {
		err := WriteMonthTemplate(ws, m)
		if err != nil {
			return err
		}
	}
	return nil
}
