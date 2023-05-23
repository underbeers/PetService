package service

import (
	"github.com/gin-gonic/gin"
	"github.com/underbeers/PetService/pkg/models"
	"github.com/underbeers/PetService/pkg/repository"
)

type PetType interface {
	GetAll(models.PetTypeFilter) ([]models.PetType, error)
}

type Breed interface {
	GetAll(filter models.BreedFilter) ([]models.Breed, error)
}

type PetCard interface {
	Create(petCard models.PetCard) error
	GetAll(filter models.PetCardFilter) ([]models.PetCard, error)
	GetMain(filter models.PetCardFilter) ([]models.PetCardMainInfo, error)
	Delete(id int) error
	Update(id int, input models.UpdateCardInput) error
	SetImage(id int, imageThumbnailLink, imageOriginLink string) error
}

type Service struct {
	PetType
	Breed
	PetCard
	Config *repository.Config
	Router *gin.Engine
}

func NewService(repos *repository.Repository, cfg *repository.Config) *Service {
	return &Service{
		PetType: NewPetTypeService(repos.PetType),
		Breed:   NewBreedService(repos.Breed),
		PetCard: NewPetCardService(repos.PetCard),
		Config:  cfg,
	}
}
