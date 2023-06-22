package date

import (
	"fmt"
	"strings"
)

type Date struct {
	day   uint
	month uint
	year  uint32
}

// Get a day from Date object
func (d *Date) Day() uint {
	return d.day
}

// Get a month from Date object
func (d *Date) Month() uint {
	return d.month
}

// Get a year from Date object
func (d *Date) Year() uint32 {
	return d.year
}

// Returns a pointer to the next date if it exists, nil otherwise
func (d *Date) Next() *Date {
	day, m, y := d.day, d.month, d.year

	err := ValidateDate(day+1, m, y)
	if err == nil {
		return &Date{day + 1, m, y}
	}
	err = ValidateDate(1, m+1, y)
	if err == nil {
		return &Date{1, m + 1, y}
	}
	// prevent year overflow, fixme
	if y == 4294967295 {
		return nil
	}
	err = ValidateDate(1, 1, y+1)
	if err == nil {
		return &Date{1, 1, y + 1}
	}

	return nil
}

// Returns a pointer to the previous date if it exists, nil otherwise
func (d *Date) Previous() *Date {
	day, m, y := d.day, d.month, d.year

	err := ValidateDate(day-1, m, y)
	if err == nil {
		return &Date{day - 1, m, y}
	}
	if m == 3 {
		if IsYearLeap(y) {
			err = ValidateDate(29, 2, y)
			if err == nil {
				return &Date{29, 2, y}
			}
		} else {
			err = ValidateDate(28, 2, y)
			if err == nil {
				return &Date{28, 2, y}
			}
		}
	} else {
		err = ValidateDate(31, m-1, y)
		if err == nil {
			return &Date{31, m - 1, y}
		}
		err = ValidateDate(30, m-1, y)
		if err == nil {
			return &Date{30, m - 1, y}
		}
	}
	// prevent year 0
	if y == 1 {
		return nil
	}
	err = ValidateDate(31, 12, y-1)
	if err == nil {
		return &Date{31, 12, y - 1}
	}

	return nil
}

// Checks if the date's year is a leap year
func (d *Date) IsYearLeap() bool {
	return IsYearLeap(d.year)
}

// Returns the date in iso format
func (d *Date) ToIso() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.year, d.month, d.day)
}

// Returns an error if the date is invalid
func ValidateDate(day, month uint, year uint32) error {
	if month > 12 || month < 1 {
		return fmt.Errorf("Month %d is not possible", month)
	}
	if day > 31 || day < 1 {
		return fmt.Errorf("Day %d is not possible", day)
	}
	if year < 1 {
		return fmt.Errorf("Year %d is not possible", year)
	}

	// February case
	if month == 2 {
		if day > 29 {
			return fmt.Errorf("There are maximum 29 days in February")
		}
		if !IsYearLeap(year) && day == 29 {
			return fmt.Errorf("There are 28 days in February in year %d", year)
		}
		return nil
	}

	if day > 30 {
		if month < 8 && month%2 == 0 || month >= 8 && month%2 == 1 {
			return fmt.Errorf("There are 30 days in %dth month", month)
		}
	}
	return nil
}

// Checks if the year is leap year
func IsYearLeap(year uint32) bool {
	if year == 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}

// Build the Date object from day, month and year
func BuildDate(day, month uint, year uint32) (*Date, error) {
	err := ValidateDate(day, month, year)
	if err != nil {
		return nil, err
	}

	d := &Date{
		day:   day,
		month: month,
		year:  year,
	}
	return d, nil
}

// Build the Date object from iso string ("year-month-day")
func BuildDateFromIso(dateIso string) (*Date, error) {
	var day, month uint
	var year uint32
	_, err := fmt.Sscanf(dateIso, "%d-%d-%d", &year, &month, &day)

	if err != nil {
		return nil, err
	}

	d, err := BuildDate(day, month, year)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Convert a month name to it's number (1-indexed)
func MonthToNumber(month string) (uint, error) {
	var monthMap = map[string]uint{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
	num, exists := monthMap[strings.ToLower(month)]
	if !exists {
		return 0, fmt.Errorf("There is no %s month", month)
	}
	return num, nil
}

// Convert a month number (1-indexed) to name
func NumberToMonth(month uint) (string, error) {
	var monthMap = map[uint]string{
		1:  "january",
		2:  "february",
		3:  "march",
		4:  "april",
		5:  "may",
		6:  "june",
		7:  "july",
		8:  "august",
		9:  "september",
		10: "october",
		11: "november",
		12: "december",
	}
	val, exists := monthMap[month]
	if !exists {
		return "", fmt.Errorf("Possible values are 1-12. You passed %d", month)
	}
	return val, nil
}
