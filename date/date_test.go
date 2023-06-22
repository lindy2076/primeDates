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

func TestNumberToMonth(t *testing.T) {
	var data = []struct {
		name     string
		input    uint
		err      bool
		expected string
	}{
		{"0", 0, true, ""},
		{"1", 1, false, "january"},
		{"2", 2, false, "february"},
		{"3", 3, false, "march"},
		{"4", 4, false, "april"},
		{"5", 5, false, "may"},
		{"12", 12, false, "december"},
		{"13", 13, true, ""},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res, err := NumberToMonth(tt.input)
			if err != nil {
				if !tt.err {
					t.Errorf("NumberToMonth returned error %s", err)
				}
			} else {
				if res != tt.expected {
					t.Errorf("NumberToMonth got %s, expected %s", res, tt.expected)
				}
			}
		})
	}
}

func TestMonthToNumber(t *testing.T) {
	var data = []struct {
		name     string
		input    string
		err      bool
		expected uint
	}{
		{"randomstr", "qwersfa", true, 0},
		{"jan", "january", false, 1},
		{"dec", "december", false, 12},
		{"none", "", true, 0},
		{"jan2", "JANUARY", false, 1},
		{"jan3", "January", false, 1},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			res, err := MonthToNumber(tt.input)
			if err != nil {
				if !tt.err {
					t.Errorf("NumberToMonth returned error %s", err)
				}
			} else {
				if res != tt.expected {
					t.Errorf("NumberToMonth got %d, expected %d", res, tt.expected)
				}
			}
		})
	}
}

func TestNext(t *testing.T) {
	var data = []struct {
		name     string
		input    [3]uint
		err      bool
		expected [3]uint
	}{
		{"30-01-2000", [3]uint{30, 1, 2000}, false, [3]uint{31, 1, 2000}},
		{"31-01-2000", [3]uint{31, 1, 2000}, false, [3]uint{1, 2, 2000}},
		{"28-02-2000", [3]uint{28, 2, 2000}, false, [3]uint{29, 2, 2000}},
		{"29-02-2000", [3]uint{29, 2, 2000}, false, [3]uint{1, 3, 2000}},
		{"28-02-2001", [3]uint{28, 2, 2001}, false, [3]uint{1, 3, 2001}},
		{"28-02-2000", [3]uint{28, 2, 2000}, false, [3]uint{29, 2, 2000}},
		{"31-12-2000", [3]uint{31, 12, 2000}, false, [3]uint{1, 1, 2001}},
		{"32-12-4294967295", [3]uint{32, 12, 4294967295}, true, [3]uint{0, 0, 0}},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{tt.input[0], uint(tt.input[1]), uint32(tt.input[2])}
			next := d.Next()
			if next == nil {
				if !tt.err {
					t.Errorf("Next got nil, expected %04d-%02d-%02d",
						tt.expected[2], tt.expected[1], tt.expected[0])
				}
			} else if next.day != tt.expected[0] || next.month != tt.expected[1] || next.year != uint32(tt.expected[2]) {
				t.Errorf("Next got %04d-%02d-%02d, expected %04d-%02d-%02d",
					next.year, next.month, next.day, tt.expected[2], tt.expected[1], tt.expected[0])
			}
		})
	}
}

func TestPrevious(t *testing.T) {
	var data = []struct {
		name     string
		input    [3]uint
		err      bool
		expected [3]uint
	}{
		{"30-01-2000", [3]uint{30, 1, 2000}, false, [3]uint{29, 1, 2000}},
		{"01-01-2000", [3]uint{1, 1, 2000}, false, [3]uint{31, 12, 1999}},
		{"01-03-2000", [3]uint{1, 3, 2000}, false, [3]uint{29, 2, 2000}},
		{"01-03-2001", [3]uint{1, 3, 2001}, false, [3]uint{28, 2, 2001}},
		{"01-04-2001", [3]uint{1, 4, 2001}, false, [3]uint{31, 3, 2001}},
		{"01-05-2000", [3]uint{1, 5, 2000}, false, [3]uint{30, 4, 2000}},
		{"01-01-0001", [3]uint{1, 1, 1}, true, [3]uint{0, 0, 0}},
		{"01-01-0002", [3]uint{1, 1, 2}, false, [3]uint{31, 12, 1}},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{tt.input[0], uint(tt.input[1]), uint32(tt.input[2])}
			next := d.Previous()
			if next == nil {
				if !tt.err {
					t.Errorf("Previous got nil, expected %04d-%02d-%02d",
						tt.expected[2], tt.expected[1], tt.expected[0])
				}
			} else if next.day != tt.expected[0] || next.month != tt.expected[1] || next.year != uint32(tt.expected[2]) {
				t.Errorf("Previous got %04d-%02d-%02d, expected %04d-%02d-%02d",
					next.year, next.month, next.day, tt.expected[2], tt.expected[1], tt.expected[0])
			}
		})
	}
}
