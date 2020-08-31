package model

type File struct {
	BaseModel  `xorm:"extends" json:"-"`
	KeyPath    string `gorm:"column:key_path"  json:"key_path"`    //秘钥
	KeyHash    string `gorm:"column:key_hash"  json:"key_hash"`    //秘钥
	ThumbPath  string `xorm:"column:thumb_path" json:"thumb_path"` //缩略图
	ThumbHash  string `xorm:"column:thumb_hash" json:"thumb_hash"`
	PosterPath string `xorm:"column:poster_path" json:"poster_path"` //海报地址
	PosterHash string `xorm:"column:poster_hash" json:"poster_hash"`
	SourcePath string `xorm:"column:source_path" json:"source_path"` //原片地址
	SourceHash string `xorm:"column:source_hash" json:"source_hash"`
	M3U8Index  string `gorm:"column:m3u8_index" json:"m3u8_index"` //M3U8名
	M3U8Path   string `xorm:"column:m3u8_path" json:"m3u8_path"`   //切片地址
	M3U8Hash   string `xorm:"column:m3u8_hash" json:"m3u8_hash"`
}

func init() {
	RegisterTable(File{})
}
