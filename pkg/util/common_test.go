package util

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name     string
		year     int64
		expected bool
	}{
		{
			name:     "2000 is leap year",
			year:     2000,
			expected: true,
		},
		{
			name:     "2020 is leap year",
			year:     2020,
			expected: true,
		},
		{
			name:     "2021 is not leap year",
			year:     2021,
			expected: false,
		},
		{
			name:     "2100 is not leap year",
			year:     2100,
			expected: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := IsLeapYear(test.year)

			if got != test.expected {
				t.Errorf("Unexpected, got: %v, expected: %v", got, test.expected)
			}
		})
	}
}
