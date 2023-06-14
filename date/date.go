package date

import "fmt"

type Date struct {
	day   uint
	month uint
	year  uint32
}

// TODO
// Returns the next date
func (d *Date) Next() *Date {
	return d
}

// TODO
// Returns the previous date
func (d *Date) Previous() *Date {
	return d
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
	_, err := fmt.Sscanf(dateIso, "%d-%d-%d", day, month, year)

	if err != nil {
		return nil, err
	}
	err = ValidateDate(day, month, year)
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
