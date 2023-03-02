package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/underbeers/PetService/pkg/models"
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
