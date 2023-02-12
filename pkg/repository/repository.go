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

type Repository struct {
	PetType
	Breed
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PetType: NewPetTypePostgres(db),
		Breed:   NewBreedPostgres(db),
	}
}
