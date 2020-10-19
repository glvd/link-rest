package model

import v1 "github.com/glvd/link-rest/restapi/v1"

type File struct {
	BaseModel  `xorm:"extends" json:"-" swaggerignore:"true"`
	RootHash   string `gorm:"column:root_hash;unique" json:"root_hash"` //跟索引
	KeyPath    string `gorm:"column:key_path"  json:"key_path"`         //秘钥
	KeyHash    string `gorm:"column:key_hash"  json:"key_hash"`         //秘钥
	ThumbPath  string `gorm:"column:thumb_path" json:"thumb_path"`      //缩略图地址
	ThumbHash  string `gorm:"column:thumb_hash" json:"thumb_hash"`      //缩略图Hash
	InfoPath   string `gorm:"column:info_path" json:"info_path"`        //媒体信息地址
	InfoHash   string `gorm:"column:info_hash" json:"info_hash"`        //媒体信息Hash
	PosterPath string `gorm:"column:poster_path" json:"poster_path"`    //海报地址
	PosterHash string `gorm:"column:poster_hash" json:"poster_hash"`    //海报Hash
	SourcePath string `gorm:"column:source_path" json:"source_path"`    //原片地址
	SourceHash string `gorm:"column:source_hash" json:"source_hash"`    //原片Hash
	M3U8Index  string `gorm:"column:m3u8_index" json:"m3u8_index"`      //M3U8名
	M3U8Path   string `gorm:"column:m3u8_path" json:"m3u8_path"`        //切片地址
	M3U8Hash   string `gorm:"column:m3u8_hash" json:"m3u8_hash"`        //切片Hash
}

func init() {
	RegisterTable(File{})
}

func (m File) TableName() string {
	return "file" + "_" + v1.Version
}
