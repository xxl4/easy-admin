package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"deletedAt"`
}
