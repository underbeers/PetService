package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/underbeers/PetService/pkg/models"
)

type BreedPostgres struct {
	db *sqlx.DB
}

func NewBreedPostgres(db *sqlx.DB) *BreedPostgres {
	return &BreedPostgres{db: db}
}

func createBreedQuery(filter models.BreedFilter) string {
	query := fmt.Sprintf(`SELECT id, pet_type_id, breed_name FROM "pet_service".public.breed`)

	if filter.BreedId != 0 {
		query += fmt.Sprintf("WHERE id = %d", filter.BreedId)
	} else if filter.PetTypeId != 0 && filter.BreedName != "" {
		query += fmt.Sprintf(`WHERE pet_type_id = %d AND breed_name = '%s'`, filter.PetTypeId, filter.BreedName)
	} else if filter.PetTypeId != 0 {
		query += fmt.Sprintf("WHERE pet_type_id = %d ORDER BY id", filter.PetTypeId)
	} else if filter.BreedName != "" {
		query += fmt.Sprintf(`WHERE breed_name = '%s'`, filter.BreedName)
	}

	return query
}

func (r *BreedPostgres) GetAll(filter models.BreedFilter) ([]models.Breed, error) {
	var breedList []models.Breed

	query := createBreedQuery(filter)
	err := r.db.Select(&breedList, query)

	return breedList, err
}
