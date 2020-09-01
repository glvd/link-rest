package model

type Media struct {
	BaseModel `json:"-"`

	InfoID string `gorm:"type:uuid;index;foreignKey:Info" json:"-"`
	Info   Info   `json:"info"`
	FileID string `gorm:"type:uuid;index;foreignKey:File" json:"-"`
	File   File   `json:"file"`
}

func init() {
	RegisterTable(Media{})
}
