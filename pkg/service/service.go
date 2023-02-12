package service

import (
	pet_service "github.com/underbeers/PetService"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetType interface {
	GetAll(pet_service.PetTypeFilter) ([]pet_service.PetType, error)
}

type Service struct {
	PetType
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		PetType: NewPetTypeService(repos.PetType),
	}
}
