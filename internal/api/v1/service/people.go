package service

import "effective_mobile_rest/internal/api/v1/repository"

type PeopleService struct {
	repo repository.People
}

func NewPeopleService(repo repository.People) *PeopleService {
	return &PeopleService{repo: repo}
}
