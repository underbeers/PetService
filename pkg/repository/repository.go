package repository

import (
	"github.com/jmoiron/sqlx"
	pet_service "github.com/underbeers/PetService"
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

type Repository struct {
	PetType
	Breed
	PetCard
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PetType: NewPetTypePostgres(db),
		Breed:   NewBreedPostgres(db),
		PetCard: NewPetCardPostgres(db),
	}
}
