package entity

import "time"

type User struct {
	FirstName     string    `json:"first_name"`
	SecondName    string    `json:"second_name"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	IncomePerYear int64     `json:"income_per_year"`
}
