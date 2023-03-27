package models

type PetType struct {
	Id   int    `json:"id" db:"id"`
	Type string `json:"petType" db:"pet_type" binding:"required"`
}
