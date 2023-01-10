package db

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt"`
}
