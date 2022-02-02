package show

import (
	"fmt"
	"mydiary/pkg/util"
)

func ValidateInput(x int64) error {
	if x < 10000000 || x >= 100000000 {
		return fmt.Errorf("input must be format like 20200101")
	}
	y := x / 10000
	m := (x % 10000) / 100
	d := x % 100
	isLeap := util.IsLeapYear(y)
	if y < 2016 || y > 2200 {
		return fmt.Errorf("Year must be between 2017 and 2200")
	}
	if m == 0 || m > 12 {
		return fmt.Errorf("Month must be between 01 and 12")
	}
	if m == 2 {
		if isLeap {
			if d == 0 || d > int64(util.DayLengths[m-1])+1 {
				return fmt.Errorf("Day must be between 01 and 29 (leap year Feburuary)")
			} else {
				return nil
			}
		} else {
			if d == 0 || d > int64(util.DayLengths[m-1]) {
				return fmt.Errorf("Day must be between 01 and 28 (not leap year Feburuary)")
			} else {
				return nil
			}
		}
	}
	if d == 0 || d > int64(util.DayLengths[m-1]) {
		return fmt.Errorf("Day must be between 01 and %d", util.DayLengths[m-1])
	}

	return nil
}
