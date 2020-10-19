package model

import v1 "github.com/glvd/link-rest/restapi/v1"

type LinkInfo struct {
	Category string `gorm:"column:category" json:"category"` //类别
}

func (LinkInfo) TableName() string {
	return "linkinfo" + "_" + v1.Version
}
