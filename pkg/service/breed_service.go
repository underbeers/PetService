package service

import (
	"github.com/underbeers/PetService/pkg/models"
	"github.com/underbeers/PetService/pkg/repository"
)

type BreedService struct {
	repo repository.Breed
}

func NewBreedService(repo repository.Breed) *BreedService {
	return &BreedService{repo: repo}
}

func (s *BreedService) GetAll(filter models.BreedFilter) ([]models.Breed, error) {
	return s.repo.GetAll(filter)
}
