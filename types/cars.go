package types

type Cars struct {
	RegNum string `db:"regNum"`
	Mark   string `db:"mark"`
	Model  string `db:"model"`
	Year   int    `db:"year"`
	People People `json:"owner"`
}

type CarsDAO struct {
	RegNum     string `json:"regNum" db:"regnum"`
	Mark       string `json:"mark" db:"mark"`
	Model      string `json:"model" db:"model"`
	Year       int    `json:"year" db:"year"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}

type CarsDTO struct {
	Id      int    `db:"id"`
	RegNum  string `db:"regNum"`
	Mark    string `db:"mark"`
	Model   string `db:"model"`
	Year    int    `db:"year"`
	OwnerId int    `db:"owner_id"`
}

type UpdateCar struct {
	RegNum *string `json:"regNum"`
	Mark   *string `json:"mark"`
	Model  *string `json:"model"`
	Year   *int    `json:"year"`
}

type CreateCar struct {
	RegNum []string `json:"regNums"`
}
