package service

import (
	"context"

	"github.com/Mitra-Apps/be-utility-service/domain/image/entity"
)

func (s *Service) GetAll(ctx context.Context) ([]*entity.Image, error) {
	images, err := s.imageRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return images, nil
}
