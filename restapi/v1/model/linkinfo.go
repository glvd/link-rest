package model

import (
	"github.com/glvd/link-rest/restapi/common/model"
	v1 "github.com/glvd/link-rest/restapi/v1"
)

type LinkInfo struct {
	Category string `gorm:"column:category" json:"category"` //类别
}

func init() {
	model.RegisterTable(LinkInfo{})
}

func (LinkInfo) TableName() string {
	return "linkinfo" + "_" + v1.Version
}
