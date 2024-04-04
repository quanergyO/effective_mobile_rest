package service

import (
	"effective_mobile_rest/internal/api/v1/repository"
	"effective_mobile_rest/types"
)

type Cars interface {
	DeleteCar(id int) error
	UpdateCarById(id int, newData types.UpdateCar) error
	GetAllCars(limit, offset int) ([]types.Cars, error)
	CreateCars(cars []types.Cars) error
}

type People interface {
}

type Service struct {
	Cars
	People
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Cars: NewCarsService(repos),
	}
}
