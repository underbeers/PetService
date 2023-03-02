package models

type Breed struct {
	Id        int    `json:"id" db:"id"`
	PetTypeId int    `json:"pet_type_id" db:"pet_type_id" binding:"required"`
	BreedName string `json:"breed_name" db:"breed_name" binding:"required"`
}
