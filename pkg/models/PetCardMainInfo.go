package models

import "time"

type PetCardMainInfo struct {
	Id          int       `json:"id" db:"id"`
	PetTypeName string    `json:"petType" db:"pet_type"`
	Name        string    `json:"petName" db:"pet_name" binding:"required"`
	Gender      string    `json:"gender" db:"gender"`
	BreedName   string    `json:"breed" db:"breed_name"`
	Photo       string    `json:"photo" db:"photo"`
	BirthDate   time.Time `json:"birthDate" db:"birth_date"`
}
