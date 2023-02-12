package service

import (
	pet_service "github.com/underbeers/PetService"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetTypeService struct {
	repo repository.PetType
}

func NewPetTypeService(repo repository.PetType) *PetTypeService {
	return &PetTypeService{repo: repo}
}

func (s *PetTypeService) GetAll(filter pet_service.PetTypeFilter) ([]pet_service.PetType, error) {
	return s.repo.GetAll(filter)
}
