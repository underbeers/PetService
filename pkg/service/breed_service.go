package service

import (
	pet_service "github.com/underbeers/PetService"
	"github.com/underbeers/PetService/pkg/repository"
)

type BreedService struct {
	repo repository.Breed
}

func NewBreedService(repo repository.Breed) *BreedService {
	return &BreedService{repo: repo}
}

func (s *BreedService) GetAll(filter pet_service.BreedFilter) ([]pet_service.Breed, error) {
	return s.repo.GetAll(filter)
}
