package model

import v1 "github.com/glvd/link-rest/restapi/v1"

type Media struct {
	BaseModel  `json:"-" swaggerignore:"true"`
	InfoID     string `gorm:"type:uuid;index;foreignKey:Info" json:"-"`
	Info       Info   `json:"info"`
	LinkInfoID string `gorm:"type:uuid;index;foreignKey:LinkInfo" json:"-"`
	LinkInfo   Info   `json:"linkinfo"`
	FileID     string `gorm:"type:uuid;index;foreignKey:File" json:"-"`
	File       File   `json:"file"`
}

func (m Media) TableName() string {
	return "media" + "_" + v1.Version
}

func init() {
	RegisterTable(Media{})
}
