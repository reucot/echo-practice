package entity

import (
	"fmt"
	"strings"
	"time"
)

type DateOfBirth struct {
	time.Time
}

type YearOfBirth struct {
	time.Time
}

func (dob DateOfBirth) MarshalJSON() ([]byte, error) {
	date := dob.Time.Format("01.02.2006")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (dob *DateOfBirth) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("01.02.2006", s)
	if err != nil {
		return err
	}

	if date.After(time.Date(time.Now().Year(), 0, 0, 0, 0, 0, 0, date.Location())) {
		return fmt.Errorf("year of date_of_birth should be before current year")
	}

	if date.Before(time.Date(1900, 0, 0, 0, 0, 0, 0, date.Location())) {
		return fmt.Errorf("year of date_of_birth should be after 1900 year")
	}

	dob.Time = date
	return
}

//Используется для биндинга query параметра
func (yob *YearOfBirth) UnmarshalParam(param string) (err error) {
	s := strings.Trim(param, "\"")

	date, err := time.Parse("2006", s)
	if err != nil {
		return err
	}

	if date.After(time.Date(time.Now().Year(), 0, 0, 0, 0, 0, 0, date.Location())) {
		return fmt.Errorf("year_of_birth should be before current year")
	}

	yob.Time = date
	return
}
