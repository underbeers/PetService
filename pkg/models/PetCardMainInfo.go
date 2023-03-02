package models

import "time"

type PetCardMainInfo struct {
	Id          int       `json:"id" db:"id"`
	PetTypeName string    `json:"pet_type" db:"pet_type"`
	Name        string    `json:"pet_name" db:"pet_name" binding:"required"`
	Gender      string    `json:"gender" db:"gender"`
	BreedName   string    `json:"breed" db:"breed_name"`
	Photo       string    `json:"photo" db:"photo"`
	BirthDate   time.Time `json:"birth_date" db:"birth_date"`
}
