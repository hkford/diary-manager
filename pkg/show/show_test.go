package show

import (
	"fmt"
	"mydiary/pkg/initialize"
	"mydiary/pkg/workspace"
	"testing"
	"time"

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
			name:     "Day invalid for not leap year Feburuary",
			input:    20210229,
			expected: Date{},
			wantErr:  true,
		},
		{
			name:     "Valid for leap year Feburuary",
			input:    20200229,
			expected: Date{2020, 2, 29},
			wantErr:  false,
		},
		{
			name:     "Valid for April",
			input:    20220421,
			expected: Date{2022, 4, 21},
			wantErr:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
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

func setupTestWorkspace() (workspace.Workspace, error) {
	fs := afero.NewMemMapFs()
	ws := workspace.Workspace{
		DiaryDir: "2020",
		IsLeap:   true,
		Fs:       &afero.Afero{Fs: fs},
	}

	err := ws.Create()
	if err != nil {
		err = fmt.Errorf("failed to create 2020 workspace: %v", err)
		return ws, err
	}
	err = initialize.WriteMonthTemplate(ws, time.January)
	if err != nil {
		err = fmt.Errorf("failed to create January template: %v", err)
		return ws, err
	}
	return ws, nil
}

func TestIsDiaryFileExists(t *testing.T) {
	ws, err := setupTestWorkspace()
	if err != nil {
		t.Errorf("TestIsDiaryFileExists failed before executing test: %v", err)
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
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
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

func TestGetDiary(t *testing.T) {
	ws, err := setupTestWorkspace()
	if err != nil {
		t.Errorf("TestGetDiary failed before executing test: %v", err)
	}
	tests := []struct {
		name     string
		date     Date
		expected string
		wantErr  bool
	}{
		{
			name:     "January 1st",
			date:     Date{2020, 1, 1},
			expected: "2020,January,01,Wed",
			wantErr:  false,
		},
		{
			name:     "January 31th",
			date:     Date{2020, 1, 31},
			expected: "2020,January,31,Fri",
			wantErr:  false,
		},
		{
			name:     "Feburuary should not exist",
			date:     Date{2020, 2, 1},
			expected: "",
			wantErr:  true,
		},
		{
			name:     "2021 should not exist",
			date:     Date{2021, 1, 1},
			expected: "",
			wantErr:  true,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := GetDiary(ws, test.date)
			if test.wantErr && err == nil {
				t.Errorf("Should raise error at %v\n", test.date)
			}
			if !test.wantErr && err != nil {
				t.Errorf("Raised error: %v", err)
			}
			if !test.wantErr && got != test.expected {
				t.Errorf("Got wrong diary, got: %v, expected: %v", got, test.expected)
			}
		})
	}
}
