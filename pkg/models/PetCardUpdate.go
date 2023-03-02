package models

import (
	"errors"
	"time"
)

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
