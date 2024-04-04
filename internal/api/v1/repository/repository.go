package repository

import (
	"effective_mobile_rest/internal/api/v1/repository/postgres"
	"effective_mobile_rest/types"
	"github.com/jmoiron/sqlx"
)

type Cars interface {
	DeleteCar(id int) error
	UpdateCarById(id int, newData types.UpdateCar) error
	GetAllCars(limit, offset int) ([]types.Cars, error)
	CreateCars(cars []types.Cars) error
}

type People interface {
}

type Repository struct {
	Cars
	People
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Cars:   postgres.NewCars(db),
		People: postgres.NewPeople(db),
	}
}
