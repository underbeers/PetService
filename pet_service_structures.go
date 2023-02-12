package PetService

import (
	"errors"
	"time"
)

type PetType struct {
	Id   int    `json:"id" db:"id"`
	Type string `json:"pet_type" db:"pet_type" binding:"required"`
}

type Breed struct {
	Id        int    `json:"id" db:"id"`
	PetTypeId int    `json:"pet_type_id" db:"pet_type_id" binding:"required"`
	BreedName string `json:"breed_name" db:"breed_name" binding:"required"`
}

type PetCard struct {
	Id            int       `json:"id" db:"id"`
	PetTypeId     int       `json:"pet_type_id" db:"pet_type_id" binding:"required"`
	PetTypeName   string    `json:"pet_type" db:"pet_type"`
	UserId        int       `json:"user_id" db:"user_id" binding:"required"`
	Name          string    `json:"pet_name" db:"pet_name" binding:"required"`
	BreedId       int       `json:"breed_id" db:"breed_id"`
	BreedName     string    `json:"breed" db:"breed_name"`
	Photo         string    `json:"photo" db:"photo"`
	BirthDate     time.Time `json:"birth_date" db:"birth_date"`
	Male          bool      `json:"male" db:"male"`
	Gender        string    `json:"gender" db:"gender"`
	Color         string    `json:"color" db:"color"`
	Care          string    `json:"care" db:"care"`
	Character     string    `json:"pet_character" db:"pet_character"`
	Pedigree      string    `json:"pedigree" db:"pedigree"`
	Sterilization bool      `json:"sterilization" db:"sterilization"`
	Vaccinations  bool      `json:"vaccinations" db:"vaccinations"`
}

type PetCardMainInfo struct {
	Id          int       `json:"id" db:"id"`
	PetTypeName string    `json:"pet_type" db:"pet_type"`
	Name        string    `json:"pet_name" db:"pet_name" binding:"required"`
	Gender      string    `json:"gender" db:"gender"`
	BreedName   string    `json:"breed" db:"breed_name"`
	Photo       string    `json:"photo" db:"photo"`
	BirthDate   time.Time `json:"birth_date" db:"birth_date"`
}

type UpdateCardInput struct {
	PetTypeId     *int       `json:"pet_type_id"`
	UserId        *int       `json:"user_id"`
	Name          *string    `json:"pet_name"`
	BreedId       *int       `json:"breed_id"`
	Photo         *string    `json:"photo"`
	BirthDate     *time.Time `json:"birth_date"`
	Male          *bool      `json:"male"`
	Color         *string    `json:"color"`
	Care          *string    `json:"care"`
	Character     *string    `json:"pet_character"`
	Pedigree      *string    `json:"pedigree"`
	Sterilization *bool      `json:"sterilization"`
	Vaccinations  *bool      `json:"vaccinations"`
}

func (i UpdateCardInput) Validate() error {
	if i.UserId == nil && i.Name == nil && i.PetTypeId == nil && i.BreedId == nil && i.Photo == nil &&
		i.BirthDate == nil && i.Male == nil && i.Color == nil && i.Care == nil && i.Character == nil &&
		i.Pedigree == nil && i.Sterilization == nil && i.Vaccinations == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
