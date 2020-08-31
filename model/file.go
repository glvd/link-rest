package model

type File struct {
	BaseModel  `xorm:"extends" json:"-"`
	ThumbPath  string `xorm:"thumb_path" json:"thumb_path"` //缩略图
	ThumbHash  string `xorm:"thumb_hash" json:"thumb_hash"`
	PosterPath string `xorm:"poster_path" json:"poster_path"` //海报地址
	PosterHash string `xorm:"poster_hash" json:"poster_hash"`
	SourcePath string `xorm:"source_path" json:"source_path"` //原片地址
	SourceHash string `xorm:"source_hash" json:"source_hash"`
	M3U8Path   string `xorm:"m3u8_path" json:"m3u8_path"` //切片地址
	M3U8Hash   string `xorm:"m3u8_hash" json:"m3u8_hash"`
}

func init() {
	RegisterTable(&File{})
}
