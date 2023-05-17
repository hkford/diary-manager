package show

import (
	"fmt"
	"mydiary/pkg/util"
	"mydiary/pkg/workspace"
	"strings"

	"github.com/spf13/afero"
)

type Date struct {
	Y int64
	M int64
	D int64
}

func ValidateInput(x int64) (date Date, err error) {
	if x < 10000000 || x >= 100000000 {
		return Date{0, 0, 0}, fmt.Errorf("input must be format like 20200101")
	}
	y := x / 10000
	m := (x % 10000) / 100
	d := x % 100
	isLeap := util.IsLeapYear(y)
	date = Date{y, m, d}
	if y < 2016 || y > 2200 {
		return date, fmt.Errorf("Year must be between 2017 and 2200")
	}
	if m == 0 || m > 12 {
		return date, fmt.Errorf("Month must be between 01 and 12")
	}
	if m == 2 {
		if isLeap {
			if d == 0 || d > int64(util.DayLengths[m-1])+1 {
				return date, fmt.Errorf("Day must be between 01 and 29 (leap year Feburuary)")
			} else {
				return date, nil
			}
		} else {
			if d == 0 || d > int64(util.DayLengths[m-1]) {
				return date, fmt.Errorf("Day must be between 01 and 28 (not leap year Feburuary)")
			} else {
				return date, nil
			}
		}
	}
	if d == 0 || d > int64(util.DayLengths[m-1]) {
		return date, fmt.Errorf("Day must be between 01 and %d", util.DayLengths[m-1])
	}

	return date, nil
}

func IsDiaryFileExists(ws workspace.Workspace, date Date) (bool, error) {
	filename := fmt.Sprintf("diaries/%v/%v%02v.txt", date.Y, date.Y, date.M)
	isDiaryExist, err := afero.Exists(ws.Fs, filename)
	if err != nil {
		return false, fmt.Errorf("Failed to run afero.Exists: *PathError %v\n", err)
	}
	return isDiaryExist, nil
}

func GetDiary(ws workspace.Workspace, date Date) (string, error) {
	filename := fmt.Sprintf("diaries/%v/%v%02v.txt", date.Y, date.Y, date.M)
	diaryOfMonth, err := afero.ReadFile(ws.Fs, filename)
	if err != nil {
		return "", fmt.Errorf("Failed to run afero.ReadFile: *PathError %v\n", err)
	}
	diaryOfDay, err := extractDiary(string(diaryOfMonth), date)
	return diaryOfDay, err
}

func extractDiary(sentences string, date Date) (string, error) {
	sentencesArray := strings.Split(sentences, "\n\n")
	dayFormat := fmt.Sprintf("%v,%v,%02v", date.Y, util.MonthNames[date.M-1], date.D)
	for _, sentence := range sentencesArray {
		if strings.HasPrefix(sentence, dayFormat) {
			return sentence, nil
		}
	}
	return "", fmt.Errorf("Could not extract diary of specified date")
}
