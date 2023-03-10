package models

import "github.com/google/uuid"

type PetTypeFilter struct {
	PetTypeId int
	PetType   string
}

type BreedFilter struct {
	BreedId   int
	PetTypeId int
	BreedName string
}

type PetCardFilter struct {
	PetCardId int
	UserId    uuid.UUID
}
