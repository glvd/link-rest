package model

type File struct {
	BaseModel  `xorm:"extends" json:"-"`
	RootHash   string `gorm:"column:root" json:"root_hash"`        //跟索引
	KeyPath    string `gorm:"column:key_path"  json:"key_path"`    //秘钥
	KeyHash    string `gorm:"column:key_hash"  json:"key_hash"`    //秘钥
	ThumbPath  string `gorm:"column:thumb_path" json:"thumb_path"` //缩略图
	ThumbHash  string `gorm:"column:thumb_hash" json:"thumb_hash"`
	PosterPath string `gorm:"column:poster_path" json:"poster_path"` //海报地址
	PosterHash string `gorm:"column:poster_hash" json:"poster_hash"`
	SourcePath string `gorm:"column:source_path" json:"source_path"` //原片地址
	SourceHash string `gorm:"column:source_hash" json:"source_hash"`
	M3U8Index  string `gorm:"column:m3u8_index" json:"m3u8_index"` //M3U8名
	M3U8Path   string `gorm:"column:m3u8_path" json:"m3u8_path"`   //切片地址
	M3U8Hash   string `gorm:"column:m3u8_hash" json:"m3u8_hash"`
}

func init() {
	RegisterTable(File{})
}
