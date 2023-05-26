package workspace

import (
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestCreate(t *testing.T) {
	fs := afero.NewMemMapFs()
	ws := Workspace{
		DiaryDir: "2020",
		IsLeap:   true,
		Fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		t.Fatal("Failed to create workspace in TestCreate")
	}
}

func TestGetDate(t *testing.T) {
	tests := []struct {
		name     string
		year     int64
		month    int64
		day      int64
		expected string
	}{
		{
			name:     "2022/01/28 is Friday",
			year:     2022,
			month:    1,
			day:      28,
			expected: "Fri",
		},
		{
			name:     "2023/05/24 is Wednesday",
			year:     2023,
			month:    5,
			day:      24,
			expected: "Wed",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := getDate(test.year, test.month, test.day)
			if got != test.expected {
				t.Errorf("Unexpected, got: %v, expected: %v", got, test.expected)
			}
		})
	}
}

func TestGenerateDayFormat(t *testing.T) {
	tests := []struct {
		name     string
		year     int64
		month    time.Month
		day      int64
		expected string
	}{
		{
			name:     "2022/01/01 is Saturday",
			year:     2022,
			month:    time.January,
			day:      1,
			expected: "2022,January,01,Sat\n\n",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := generateDayFormat(test.year, test.month, test.day)
			if got != test.expected {
				t.Errorf("Unexpected, got: %v, expected: %v", got, test.expected)
			}
		})
	}
}

func TestWriteYearTemplates(t *testing.T) {
	fs := afero.NewMemMapFs()
	ws := Workspace{
		DiaryDir: "2020",
		IsLeap:   true,
		Fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		t.Fatal("Failed to create workspace in TestWriteMonthTemplate")
	}
	err = ws.WriteYearTemplates()
	if err != nil {
		t.Fatal("Failed to write 2020 diary template in TestWriteMonthTemplate")
	}
	_, err = fs.Open("diaries/2020/202001.txt")
	if err != nil {
		t.Errorf("Failed to open January diary template in TestWriteMonthTemplate\n")
	}
}
