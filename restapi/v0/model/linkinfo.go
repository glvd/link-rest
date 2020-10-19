package model

type LinkInfo struct {
	Category string `gorm:"column:category" json:"category"` //类别
}
