package service

import (
	"github.com/underbeers/PetService/pkg/models"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetTypeService struct {
	repo repository.PetType
}

func NewPetTypeService(repo repository.PetType) *PetTypeService {
	return &PetTypeService{repo: repo}
}

func (s *PetTypeService) GetAll(filter models.PetTypeFilter) ([]models.PetType, error) {
	return s.repo.GetAll(filter)
}
