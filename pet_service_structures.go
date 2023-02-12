package PetService

import "time"

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
	Color         string    `json:"color" db:"color"`
	Care          string    `json:"care" db:"care"`
	Character     string    `json:"pet_character" db:"pet_character"`
	Pedigree      string    `json:"pedigree" db:"pedigree"`
	Sterilization bool      `json:"sterilization" db:"sterilization"`
	Vaccinations  bool      `json:"vaccinations" db:"vaccinations"`
}
