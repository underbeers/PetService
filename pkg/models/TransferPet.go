package models

import "github.com/google/uuid"

type TransferPet struct {
	Id         int       `json:"petCardID"`
	NewOwnerID uuid.UUID `json:"newOwnerID"`
}
