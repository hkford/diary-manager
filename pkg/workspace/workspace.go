package workspace

import (
	"fmt"
	"math"
	"mydiary/pkg/util"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/spf13/afero"
)

type Workspace struct {
	DiaryDir string
	IsLeap   bool
	Fs       *afero.Afero
}

func generateDayFormat(y int64, m time.Month, d int64) string {
	f := fmt.Sprintf("%v,%v,%02v,%v\n\n", y, m, d, getDate(y, int64(m), d))
	return f
}

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

// Create diary directory if it does not already exist.
func (ws *Workspace) Create() error {
	dirPath := fmt.Sprintf("diaries/%v", ws.DiaryDir)
	err := ws.Fs.MkdirAll(dirPath, 0755)
	return err
}

func (ws *Workspace) writeMonthTemplate(month time.Month) error {
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
		dayFormat := generateDayFormat(year, month, int64(d))
		template = append(template, []byte(dayFormat)...)
	}
	filename := fmt.Sprintf("diaries/%v/%v%02v.txt", ws.DiaryDir, ws.DiaryDir, strconv.FormatInt(int64(month), 10))
	err = ws.Fs.WriteFile(filename, template, 0755)
	return err
}

func (ws *Workspace) WriteYearTemplates() error {
	var g errgroup.Group
	for m := time.January; m <= time.December; m++ {
		g.Go(func() error {
			return ws.writeMonthTemplate(m)
		})
		if err := g.Wait(); err != nil {
			return err
		}
	}
	return nil
}

// Delete diary directory if generating yearly template failed.
func (ws *Workspace) Delete() error {
	dirPath := fmt.Sprintf("diaries/%v", ws.DiaryDir)
	err := ws.Fs.RemoveAll(dirPath)
	return err
}
