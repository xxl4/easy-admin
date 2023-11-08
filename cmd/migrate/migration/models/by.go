package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time             `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time             `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index;comment:flag"`
}
