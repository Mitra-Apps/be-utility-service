package service

import (
	"context"

	"github.com/Mitra-Apps/be-utility-service/domain/image/entity"
	"github.com/Mitra-Apps/be-utility-service/domain/image/repository"
)

type Service struct {
	imageRepository repository.ImageInterface
}

func New(imageRepository repository.ImageInterface) *Service {
	return &Service{imageRepository: imageRepository}
}

type ServiceInterface interface {
	GetAll(ctx context.Context) ([]*entity.Image, error)
}
