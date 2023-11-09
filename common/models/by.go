package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// create and update by user
type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:create user"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:update user"`
}

// SetCreateBy creater id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy edit id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:key"`
}

// Org model Time and delete flag
type ModelTime struct {
	CreatedAt time.Time             `json:"createdAt" gorm:"comment:create date"`
	UpdatedAt time.Time             `json:"updatedAt" gorm:"comment:update date"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index;comment:flag"`
}
