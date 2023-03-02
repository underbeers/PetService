package models

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
	UserId    int
}
