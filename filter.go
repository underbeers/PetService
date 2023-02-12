package PetService

type PetTypeFilter struct {
	PetTypeId int
	PetType   string
}

type BreedFilter struct {
	BreedId   int
	PetTypeId int
	BreedName string
}
