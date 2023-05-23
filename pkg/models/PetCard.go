package models

import (
	"github.com/google/uuid"
	"time"
)

type PetCard struct {
	Id             int       `json:"id" db:"id"`
	PetTypeId      int       `json:"petTypeID" db:"pet_type_id" binding:"required"`
	PetTypeName    string    `json:"petType" db:"pet_type"`
	UserId         uuid.UUID `json:"userID" db:"user_id"`
	Name           string    `json:"petName" db:"pet_name" binding:"required"`
	BreedId        int       `json:"breedID" db:"breed_id" binding:"required"`
	BreedName      string    `json:"breed" db:"breed_name"`
	Photo          string    `json:"photo" db:"origin_photo"`
	ThumbnailPhoto string    `json:"thumbnailPhoto" db:"thumbnail_photo"`
	BirthDate      time.Time `json:"birthDate" db:"birth_date" binding:"required"`
	Male           bool      `json:"male" db:"male"`
	Gender         string    `json:"gender" db:"gender"`
	Color          string    `json:"color" db:"color"`
	Care           string    `json:"care" db:"care"`
	Character      string    `json:"petCharacter" db:"pet_character"`
	Pedigree       string    `json:"pedigree" db:"pedigree"`
	Sterilization  bool      `json:"sterilization" db:"sterilization"`
	Vaccinations   bool      `json:"vaccinations" db:"vaccinations"`
}
