package models

type PetType struct {
	Id   int    `json:"id" db:"id"`
	Type string `json:"pet_type" db:"pet_type" binding:"required"`
}
