package entity

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
