package initialize

import "testing"

func TestIsLeapYear(t *testing.T) {
	var result bool
	result = IsLeapYear(2000)
	if result != true {
		t.Fatal("failed IsLeapYear year=2000")
	}
	result = IsLeapYear(2020)
	if result != true {
		t.Fatal("failed IsLeapYear year=2020")
	}
	result = IsLeapYear(2021)
	if result != false {
		t.Fatal("failed IsLeapYear year=2021")
	}
	result = IsLeapYear(2100)
	if result != false {
		t.Fatal("failed IsLeapYear year=2100")
	}
}

func TestGetDate(t *testing.T) {
	var result string
	result = GetDate(2022, 1, 28)
	if result != "Fri" {
		t.Fatal("failed GetDate 2022/01/28")
	}
	result = GetDate(2020, 3, 9)
	if result != "Mon" {
		t.Fatal("failed GetDate 2020/03/09")
	}
}

func TestGenerateDayFormat(t *testing.T) {
	var result string
	result = GenerateDayFormat(2022, 1, 1)
	if result != "2022,January,01,Sat\n\n" {
		t.Fatal("failed GenerateDayformat(2022,1,1)")
	}
}
