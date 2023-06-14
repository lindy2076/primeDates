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

// TODO
// Returns an error if the date is invalid
func ValidateDate(day, month uint, year uint32) error {
	if month > 12 || month < 1 {
		return fmt.Errorf("Month %d is not possible", month)
	}
	if day == 0 {
		return fmt.Errorf("Day 0 is not possible")
	}

	return nil
}

// FIXME validate data
// Build the Date object from day, month and year
func BuildDate(day, month uint, year uint32) (*Date, error) {
	d := &Date{
		day:   day,
		month: month,
		year:  year,
	}
	return d, nil
}

// FIXME validate date
// Build the Date object from iso string ("year-month-day")
func BuildDateFromIso(dateIso string) (*Date, error) {
	var day, month uint
	var year uint32
	_, err := fmt.Sscanf(dateIso, "%d-%d-%d", day, month, year)

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
