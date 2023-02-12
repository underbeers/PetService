package service

import (
	pet_service "github.com/underbeers/PetService"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetType interface {
	GetAll(pet_service.PetTypeFilter) ([]pet_service.PetType, error)
}

type Breed interface {
	GetAll(filter pet_service.BreedFilter) ([]pet_service.Breed, error)
}

type Service struct {
	PetType
	Breed
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		PetType: NewPetTypeService(repos.PetType),
		Breed:   NewBreedService(repos.Breed),
	}
}
