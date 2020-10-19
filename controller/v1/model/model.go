package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type StringArray []string

type BaseModel struct {
	ID        string         `gorm:"id primarykey default:uuid_generate_v3() comment(默认主键)" swaggerignore:"true"`
	CreatedAt time.Time      `gorm:"created_at comment(创建时间)" swaggerignore:"true"`
	UpdatedAt time.Time      `gorm:"updated_at comment(更新时间)" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `gorm:"updated_at index comment(删除时间)" swaggerignore:"true"`
	Version   int            `gorm:"version comment(版本)" swaggerignore:"true"`
}

//func (model *BaseModel) BeforeInsert() {
//	if model.ID == "" {
//		model.ID = uuid.Must(uuid.NewRandom()).String()
//	}
//}

func (model *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if model.ID == "" {
		model.ID = uuid.Must(uuid.NewRandom()).String()
	}
	return
}

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, s)
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (s StringArray) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *StringArray) GormDataType() string {
	return "json"
}
