package models

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type UpdateCardInput struct {
	PetTypeId     *int       `json:"petTypeID"`
	UserId        *uuid.UUID `json:"userID"`
	Name          *string    `json:"petName"`
	BreedId       *int       `json:"breedID"`
	Photo         *string    `json:"photo"`
	BirthDate     *time.Time `json:"birthDate"`
	Male          *bool      `json:"male"`
	Color         *string    `json:"color"`
	Care          *string    `json:"care"`
	Character     *string    `json:"petCharacter"`
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
