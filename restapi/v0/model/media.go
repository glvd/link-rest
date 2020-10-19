package model

import "github.com/glvd/link-rest/restapi/common/model"

type Media struct {
	model.BaseModel `json:"-" swaggerignore:"true"`
	InfoID          string `gorm:"type:uuid;index;foreignKey:Info" json:"-"`
	Info            Info   `json:"info"`
	FileID          string `gorm:"type:uuid;index;foreignKey:File" json:"-"`
	File            File   `json:"file"`
}

func init() {
	model.RegisterTable(Media{})
}
