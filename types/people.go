package types

type People struct {
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}

type PeopleDTO struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Surname    string `db:"surname"`
	Patronymic string `db:"patronymic"`
}
