package model

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        string     `xorm:"uuid pk comment(默认主键)"`
	CreatedAt time.Time  `xorm:"created comment(创建时间) 'created_at'"`
	UpdatedAt time.Time  `xorm:"updated comment(更新时间) 'updated_at'" `
	DeletedAt *time.Time `xorm:"deleted comment(删除时间) 'deleted_at'"`
	Version   int        `xorm:"version comment(版本)"`
}

func (model *BaseModel) BeforeInsert() {
	if model.ID == "" {
		model.ID = uuid.Must(uuid.NewRandom()).String()
	}
}