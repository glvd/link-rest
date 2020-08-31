package model

type Info struct {
	BaseModel `xorm:"extends" json:"-"`
	Thumb     Hash `xorm:"thumb_hash" json:"thumb_hash"`   //缩略图
	Poster    Hash `xorm:"poster_hash" json:"poster_hash"` //海报地址
	Source    Hash `xorm:"source_hash" json:"source_hash"` //原片地址
	M3U8      Hash `xorm:"m3u8_hash" json:"m3u8_hash"`     //切片地址
}

func init() {
	RegisterTable(&Info{})
}
