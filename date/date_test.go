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
		{"01-00-2000", [3]int{1, 0, 2000}, true, "Month 0 is not possible"},
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
