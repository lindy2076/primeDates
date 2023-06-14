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
