package service

import (
	"effective_mobile_rest/internal/api/v1/repository"
	"effective_mobile_rest/types"
)

type CarsService struct {
	repo repository.Cars
}

func NewCarsService(repo repository.Cars) *CarsService {
	return &CarsService{repo: repo}
}

func (s *CarsService) DeleteCar(id int) error {
	return s.repo.DeleteCar(id)
}

func (s *CarsService) UpdateCarById(id int, newData types.UpdateCar) error {
	return s.repo.UpdateCarById(id, newData)
}

func (s *CarsService) GetAllCars(limit, offset int) ([]types.Cars, error) {
	return s.repo.GetAllCars(limit, offset)
}

func (s *CarsService) CreateCars(cars []types.Cars) error {
	return s.repo.CreateCars(cars)
}
