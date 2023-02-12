package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	pet_service "github.com/underbeers/PetService"
)

type PetTypePostgres struct {
	db *sqlx.DB
}

func NewPetTypePostgres(db *sqlx.DB) *PetTypePostgres {
	return &PetTypePostgres{db: db}
}

func createPetTypeQuery(filter pet_service.PetTypeFilter) string {
	query := fmt.Sprintf("SELECT id, pet_type FROM %s ", petTypeTable)

	if filter.PetTypeId != 0 {
		query += fmt.Sprintf("WHERE id = %d ORDER BY id", filter.PetTypeId)
	} else if filter.PetType != "" {
		query += fmt.Sprintf(`WHERE pet_type = '%s'`, filter.PetType)
	}

	return query
}

func (r *PetTypePostgres) GetAll(filter pet_service.PetTypeFilter) ([]pet_service.PetType, error) {
	var petTypeList []pet_service.PetType

	query := createPetTypeQuery(filter)
	err := r.db.Select(&petTypeList, query)

	return petTypeList, err
}
