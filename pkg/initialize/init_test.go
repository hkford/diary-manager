package initialize

import "testing"

func TestIsLeapYear(t *testing.T) {
	var result bool
	result = IsLeapYear(2000)
	if result != true {
		t.Fatal("failed test year=2000")
	}
	result = IsLeapYear(2020)
	if result != true {
		t.Fatal("failed test year=2020")
	}
	result = IsLeapYear(2021)
	if result != false {
		t.Fatal("failed test year=2021")
	}
	result = IsLeapYear(2100)
	if result != false {
		t.Fatal("failed test year=2100")
	}
}
