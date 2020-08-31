package model

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type StringArray []string

type BaseModel struct {
	gorm.Model
	ID        string         `gorm:"id primarykey comment(默认主键)"`
	CreatedAt time.Time      `gorm:"created_at comment(创建时间)"`
	UpdatedAt time.Time      `gorm:"updated_at comment(更新时间)" `
	DeletedAt gorm.DeletedAt `gorm:"updated_at index comment(删除时间)"`
	Version   int            `gorm:"version comment(版本)"`
}

func (model *BaseModel) BeforeInsert() {
	if model.ID == "" {
		model.ID = uuid.Must(uuid.NewRandom()).String()
	}
}
