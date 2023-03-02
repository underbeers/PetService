package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/underbeers/PetService/pkg/models"
)

type PetTypePostgres struct {
	db *sqlx.DB
}

func NewPetTypePostgres(db *sqlx.DB) *PetTypePostgres {
	return &PetTypePostgres{db: db}
}

func createPetTypeQuery(filter models.PetTypeFilter) string {
	query := fmt.Sprintf("SELECT id, pet_type FROM %s ", petTypeTable)

	if filter.PetTypeId != 0 {
		query += fmt.Sprintf("WHERE id = %d ORDER BY id", filter.PetTypeId)
	} else if filter.PetType != "" {
		query += fmt.Sprintf(`WHERE pet_type = '%s'`, filter.PetType)
	}

	return query
}

func (r *PetTypePostgres) GetAll(filter models.PetTypeFilter) ([]models.PetType, error) {
	var petTypeList []models.PetType

	query := createPetTypeQuery(filter)
	err := r.db.Select(&petTypeList, query)

	return petTypeList, err
}
