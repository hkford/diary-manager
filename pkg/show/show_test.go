package show

import (
	"mydiary/pkg/initialize"
	"mydiary/pkg/util"
	"mydiary/pkg/workspace"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected Date
		wantErr  bool
	}{
		{
			name:     "Input too small",
			input:    2000000,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Input too large",
			input:    200000000,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Month and Day invalid",
			input:    20160000,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Too large Day length",
			input:    20170140,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Too large Month",
			input:    20171300,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Day invalid",
			input:    20170100,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Day invalid for January",
			input:    20170132,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Day invalid for Feburuary",
			input:    20200230,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Valid",
			input:    20200229,
			expected: Date{2020, 2, 29},
			wantErr:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ValidateInput(test.input)

			if test.wantErr && (err == nil) {
				t.Errorf("Error expected, got: %v but err is nil", got)
			}

			if !test.wantErr && !cmp.Equal(got, test.expected) {
				t.Errorf("Unexpected date, got: %v, expected: %v", got, test.expected)
			}
		})
	}
}

func TestIsDiaryFileExists(t *testing.T) {
	fs := afero.NewMemMapFs()
	ws := workspace.Workspace{
		DiaryDir: "2020",
		IsLeap:   true,
		Fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		t.Fatal("failed Create workspace")
	}
	err = initialize.WriteMonthTemplate(ws, int64(1))
	if err != nil {
		t.Fatal("failed Create month template")
	}
	tests := []struct {
		name     string
		date     Date
		expected bool
	}{
		{
			name:     "January should exist",
			date:     Date{2020, 1, 29},
			expected: true,
		},
		{
			name:     "March should not exist",
			date:     Date{2020, 3, 1},
			expected: false,
		},
		{
			name:     "2021 should not exist",
			date:     Date{2021, 3, 1},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := IsDiaryFileExists(ws, test.date)

			if err != nil {
				t.Errorf("IsDiaryFileExists raised error: %v", err)
			}
			if got != test.expected {
				t.Errorf("Unexpected, got: %v, expected: %v", got, test.expected)
			}
		})
	}
}

// Write test diary file of 2020/01
func writeTestDiary(ws workspace.Workspace, t *testing.T) {
	var template = make([]byte, 0, 700)
	days := util.DayLengths[0]
	diaryOfJanuary1st := "2020,January,01,Wed\nFirst diary.\n\n"
	template = append(template, []byte(diaryOfJanuary1st)...)
	for d := 2; d <= days; d++ {
		dayFormat := initialize.GenerateDayFormat(2020, 1, int64(d))
		template = append(template, []byte(dayFormat)...)
	}
	filename := "diaries/2020/202001.txt"
	err := ws.Fs.WriteFile(filename, template, 0755)
	if err != nil {
		t.Fatal("failed to generate test diary")
	}
}

func TestGetDiary(t *testing.T) {
	var result string
	fs := afero.NewMemMapFs()
	ws := workspace.Workspace{
		DiaryDir: "2020",
		IsLeap:   true,
		Fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		t.Fatal("failed Create workspace")
	}
	writeTestDiary(ws, t)
	date := Date{2020, 1, 1}
	result, _ = GetDiary(ws, date)
	if result != "2020,January,01,Wed\nFirst diary." {
		t.Errorf("Got wrong diary: %v", result)
	}
}
