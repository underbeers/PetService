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

type PetCard interface {
	Create(petCard pet_service.PetCard) error
	GetAll(filter pet_service.PetCardFilter) ([]pet_service.PetCard, error)
	GetMain(filter pet_service.PetCardFilter) ([]pet_service.PetCardMainInfo, error)
	Delete(id int) error
	Update(id int, input pet_service.UpdateCardInput) error
}

type Service struct {
	PetType
	Breed
	PetCard
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		PetType: NewPetTypeService(repos.PetType),
		Breed:   NewBreedService(repos.Breed),
		PetCard: NewPetCardService(repos.PetCard),
	}
}
