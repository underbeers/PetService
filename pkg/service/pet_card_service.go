package service

import (
	pet_service "github.com/underbeers/PetService"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetCardService struct {
	repo repository.PetCard
}

func NewPetCardService(repo repository.PetCard) *PetCardService {
	return &PetCardService{repo: repo}
}

func (s *PetCardService) Create(petCard pet_service.PetCard) error {
	return s.repo.Create(petCard)
}

func (s *PetCardService) GetAll(filter pet_service.PetCardFilter) ([]pet_service.PetCard, error) {
	return s.repo.GetAll(filter)
}

func (s *PetCardService) GetMain(filter pet_service.PetCardFilter) ([]pet_service.PetCardMainInfo, error) {
	return s.repo.GetMain(filter)
}

func (s *PetCardService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PetCardService) Update(id int, input pet_service.UpdateCardInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
