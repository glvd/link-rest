package model

type Media struct {
	BaseModel  `json:"-" swaggerignore:"true"`
	InfoID     string `gorm:"type:uuid;index;foreignKey:Info" json:"-"`
	Info       Info   `json:"info"`
	LinkInfoID string `gorm:"type:uuid;index;foreignKey:LinkInfo" json:"-"`
	LinkInfo   Info   `json:"linkinfo"`
	FileID     string `gorm:"type:uuid;index;foreignKey:File" json:"-"`
	File       File   `json:"file"`
}

func init() {
	RegisterTable(Media{})
}
