package repository

import (
	"github.com/jmoiron/sqlx"
	pet_service "github.com/underbeers/PetService"
)

type PetType interface {
	GetAll(pet_service.PetTypeFilter) ([]pet_service.PetType, error)
}

type Repository struct {
	PetType
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PetType: NewPetTypePostgres(db),
	}
}
