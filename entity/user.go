package entity

import (
	"fmt"
	"time"
)

//Model
type User struct {
	Id            int           `json:"id"`
	FirstName     string        `json:"first_name"`
	SecondName    string        `json:"second_name"`
	DateOfBirth   DateOfBirth   `json:"date_of_birth"`
	IncomePerYear IncomePerYear `json:"income_per_year"`
}

//Filter
type FilterUser struct {
	FirstName   string      `json:"first_name" query:"first_name"`
	SecondName  string      `json:"second_name" query:"second_name"`
	YearOfBirth YearOfBirth `json:"year_of_birth" query:"year_of_birth"`
}

func (fu FilterUser) Filter() string {
	var fields []FilterFields

	if fu.FirstName != "" {
		fields = append(fields, *NewFilterField("users", "first_name", "=", fu.FirstName))
	}

	if fu.SecondName != "" {
		fields = append(fields, *NewFilterField("users", "second_name", "=", fu.SecondName))
	}

	if fu.YearOfBirth.Time != (time.Time{}) {
		fields = append(fields, *NewFilterField("extract ( year from users", "date_of_birth)", "=", fmt.Sprintf("%d", fu.YearOfBirth.Year())))
	}

	return NewFilterGroup(fields, "AND ").Filter()
}
