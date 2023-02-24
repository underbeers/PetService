package service

import (
	"github.com/underbeers/PetService/pkg/models"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetType interface {
	GetAll(models.PetTypeFilter) ([]models.PetType, error)
}

type Breed interface {
	GetAll(filter models.BreedFilter) ([]models.Breed, error)
}

type PetCard interface {
	Create(petCard models.PetCard) error
	GetAll(filter models.PetCardFilter) ([]models.PetCard, error)
	GetMain(filter models.PetCardFilter) ([]models.PetCardMainInfo, error)
	Delete(id int) error
	Update(id int, input models.UpdateCardInput) error
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
