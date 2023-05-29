package service

import (
	"github.com/underbeers/PetService/pkg/models"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetCardService struct {
	repo repository.PetCard
}

func NewPetCardService(repo repository.PetCard) *PetCardService {
	return &PetCardService{repo: repo}
}

func (s *PetCardService) Create(petCard models.PetCard) error {
	return s.repo.Create(petCard)
}

func (s *PetCardService) GetAll(filter models.PetCardFilter) ([]models.PetCard, error) {
	return s.repo.GetAll(filter)
}

func (s *PetCardService) GetMain(filter models.PetCardFilter) ([]models.PetCardMainInfo, error) {
	return s.repo.GetMain(filter)
}

func (s *PetCardService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PetCardService) Update(id int, input models.UpdateCardInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}

func (s *PetCardService) SetImage(id int, imageThumbnailLink, imageOriginLink string) error {
	return s.repo.SetImage(id, imageThumbnailLink, imageOriginLink)
}

func (s *PetCardService) TransferPet(input models.TransferPet) error {
	return s.repo.TransferPet(input)
}
