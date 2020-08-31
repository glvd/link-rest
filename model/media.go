package model

type Media struct {
	BaseModel `xorm:"extends" json:"-"`
	InfoID    string `xorm:"info_id" json:"-"`
	Info      Info   `xorm:"-" json:"info"`
	FileID    string `xorm:"file_id" json:"-"`
	File      File   `xorm:"-" json:"file"`
}

func init() {
	RegisterTable(&Media{})
}
