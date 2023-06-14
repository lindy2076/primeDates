package date

import (
	"fmt"
	"testing"
)

func TestValidateDate(t *testing.T) {
	var data = []struct {
		name   string
		input  [3]int
		err    bool
		errMsg string
	}{
		{"17-01-2000", [3]int{17, 1, 2000}, false, ""},
		{"30-04-2000", [3]int{17, 1, 2000}, false, ""},
		{"31-04-2000", [3]int{31, 4, 2000}, true, "There are 30 days in 4th month"},
		{"31-09-2000", [3]int{31, 9, 2000}, true, "There are 30 days in 9th month"},
		{"17-13-2000", [3]int{17, 13, 2000}, true, "Month 13 is not possible"},
		{"00-04-2000", [3]int{0, 1, 2000}, true, "Day 0 is not possible"},
		{"32-04-2000", [3]int{32, 1, 2000}, true, "Day 32 is not possible"},
		{"01-00-2000", [3]int{1, 0, 2000}, true, "Month 0 is not possible"},
		{"01-01-0000", [3]int{1, 1, 0}, true, "Year 0 is not possible"},
		{"28-02-2001", [3]int{28, 2, 2001}, false, ""},
		{"30-02-2001", [3]int{30, 2, 2001}, true, "There are maximum 29 days in February"},
		{"29-02-2001", [3]int{29, 2, 2001}, true, "There are 28 days in February in year 2001"},
		{"29-02-2004", [3]int{29, 2, 2004}, false, ""},
		{"29-02-2000", [3]int{29, 2, 2000}, false, ""},
		{"29-02-1900", [3]int{29, 2, 1900}, true, "There are 28 days in February in year 1900"},
		{"31-01-1900", [3]int{31, 1, 1900}, false, ""},
		{"31-03-1900", [3]int{31, 3, 1900}, false, ""},
		{"31-07-1900", [3]int{31, 7, 1900}, false, ""},
		{"31-08-1900", [3]int{31, 8, 1900}, false, ""},
		{"31-10-1900", [3]int{31, 10, 1900}, false, ""},
		{"31-04-1900", [3]int{31, 4, 1900}, true, "There are 30 days in 4th month"},
		{"31-06-1900", [3]int{31, 6, 1900}, true, "There are 30 days in 6th month"},
		{"31-09-1900", [3]int{31, 9, 1900}, true, "There are 30 days in 9th month"},
		{"31-11-1900", [3]int{31, 11, 1900}, true, "There are 30 days in 11th month"},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res := ValidateDate(uint(tt.input[0]), uint(tt.input[1]), uint32(tt.input[2]))
			if res == nil && tt.err || res != nil && !tt.err {
				t.Errorf("ValidateDate got %s, expected %s", res, tt.errMsg)
			} else if res != nil && tt.errMsg != fmt.Sprintf("%s", res) {
				// check if errMsg is different
				t.Errorf("ValidateDate got %s, expected %s", res, tt.errMsg)
			}
		})
	}
}

func TestBuildDate(t *testing.T) {
	var data = []struct {
		name   string
		input  [3]int
		err    bool
		errMsg string
	}{
		{"17-02-2000", [3]int{17, 2, 2000}, false, ""},
		{"31-02-2000", [3]int{31, 2, 2000}, true, "There are maximum 29 days in February"},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res, err := BuildDate(uint(tt.input[0]), uint(tt.input[1]), uint32(tt.input[2]))
			if err != nil && !tt.err {
				t.Errorf("BuildDate %d, %s", tt.input, err)
			}
			if err == nil && tt.err {
				t.Errorf("BuildDate %d; expected %s, got nil", tt.input, err)
			}
			if err == nil {
				if fmt.Sprintf("%T", res) != "*date.Date" {
					t.Errorf("BuildDate returned type is %T, not *Date", res)
				}
				if res.day != uint(tt.input[0]) {
					t.Errorf("BuildDate day %d, expected %d", res.day, tt.input[0])
				}
				if res.month != uint(tt.input[1]) {
					t.Errorf("BuildDate month %d, expected %d", res.day, tt.input[0])
				}
				if res.year != uint32(tt.input[2]) {
					t.Errorf("BuildDate year %d, expected %d", res.day, tt.input[0])
				}
			}
		})
	}
}

func TestBuildDateFromIso(t *testing.T) {
	var data = []struct {
		name   string
		input  string
		err    bool
		errMsg string
	}{
		{"17-02-2000", "2000-02-17", false, ""},
		{"31-02-2000", "2000-02-31", true, "There are maximum 29 days in February"},
		{"31-13-2000", "2000-13-31", true, "Month 13 is not possible"},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res, err := BuildDateFromIso(tt.input)
			if err != nil && !tt.err {
				t.Errorf("BuildDateFromIso %s, %s", tt.input, err)
			}
			if err == nil && tt.err {
				t.Errorf("BuildDateFromIso %s; expected %s, got nil", tt.input, err)
			}
			if err == nil {
				if fmt.Sprintf("%T", res) != "*date.Date" {
					t.Errorf("BuildDateFromIso returned type is %T, not *Date", res)
				}
				formattedResult := fmt.Sprintf("%04d-%02d-%02d", res.year, res.month, res.day)
				if formattedResult != tt.input {
					t.Errorf("BuildDateFromIso returned %s, not %s", formattedResult, tt.input)
				}
			}
		})
	}
}

func TestToIso(t *testing.T) {
	var data = []struct {
		name     string
		input    *Date
		expected string
	}{
		{"17-02-2000", &Date{17, 2, 2000}, "2000-02-17"},
		{"01-02-2000", &Date{1, 2, 2000}, "2000-02-01"},
		{"01-12-2000", &Date{1, 12, 2000}, "2000-12-01"},
		{"01-12-0999", &Date{1, 12, 999}, "0999-12-01"},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.input.ToIso()
			if res != tt.expected {
				t.Errorf("ToIso got %s, expected %s", res, tt.expected)
			}
		})
	}
}
