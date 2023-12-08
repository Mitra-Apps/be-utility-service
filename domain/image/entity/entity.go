package entity

import (
	"time"

	pb "github.com/Mitra-Apps/be-utility-service/domain/proto/image"
	"github.com/google/uuid"
)

type Image struct {
	Id        uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ImagePath string        `gorm:"type:varchar(255);not null;unique"`
	IsActive  bool          `gorm:"type:bool;not null;default:TRUE"`
	CreatedAt time.Time     `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy uuid.UUID     `gorm:"type:uuid;not null"`
	UpdatedAt *time.Time    `gorm:"type:timestamptz;null"`
	UpdatedBy uuid.NullUUID `gorm:"type:uuid;null"`
}

func (u *Image) ToProto() *pb.Image {
	return &pb.Image{
		Id:        u.Id.String(),
		ImagePath: u.ImagePath,
	}
}
