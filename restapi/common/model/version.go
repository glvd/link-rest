package model

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	V string
}

func init() {
	RegisterTable(Version{})
}
