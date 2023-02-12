package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	pet_service "github.com/underbeers/PetService"
)

type BreedPostgres struct {
	db *sqlx.DB
}

func NewBreedPostgres(db *sqlx.DB) *BreedPostgres {
	return &BreedPostgres{db: db}
}

func createBreedQuery(filter pet_service.BreedFilter) string {
	query := fmt.Sprintf("SELECT id, pet_type_id, breed_name FROM %s ", breedTable)

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

func (r *BreedPostgres) GetAll(filter pet_service.BreedFilter) ([]pet_service.Breed, error) {
	var breedList []pet_service.Breed

	query := createBreedQuery(filter)
	err := r.db.Select(&breedList, query)

	return breedList, err
}
