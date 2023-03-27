package models

type Breed struct {
	Id        int    `json:"id" db:"id"`
	PetTypeId int    `json:"petTypeId" db:"pet_type_id" binding:"required"`
	BreedName string `json:"breedName" db:"breed_name" binding:"required"`
}
