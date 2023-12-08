package repository

import (
	"context"

	"github.com/Mitra-Apps/be-utility-service/domain/image/entity"
)

type ImageInterface interface {
	GetAll(ctx context.Context) ([]*entity.Image, error)
}
