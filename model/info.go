package model

type Info struct {
	BaseModel    `json:"-"`
	VideoNo      string      `gorm:"video_no" json:"video_no"`           //编号
	Intro        string      `gorm:"varchar(2048)" json:"intro"`         //简介
	Alias        StringArray `gorm:"alias" json:"alias"`                 //别名，片名
	Key          string      `gorm:"key"  json:"key"`                    //秘钥
	M3U8         string      `gorm:"m3u8" json:"m3u8"`                   //M3U8名
	Role         StringArray `gorm:"role" json:"role"`                   //主演
	Director     string      `gorm:"director" json:"director"`           //导演
	Systematics  string      `gorm:"systematics" json:"systematics"`     //分级
	Season       string      `gorm:"season" json:"season"`               //季
	TotalEpisode string      `gorm:"total_episode" json:"total_episode"` //总集数
	Episode      string      `gorm:"episode" json:"episode"`             //集数
	Producer     string      `gorm:"producer" json:"producer"`           //生产商
	Publisher    string      `gorm:"publisher" json:"publisher"`         //发行商
	MediaType    string      `gorm:"media_type" json:"media_type"`       //类型：film，FanDrama
	Format       string      `gorm:"format" json:"format"`               //输出格式：3D，2D,VR(VR格式：Half-SBS：左右半宽,Half-OU：上下半高,SBS：左右全宽)
	Language     string      `gorm:"language" json:"language"`           //语言
	Caption      string      `gorm:"caption" json:"caption"`             //字幕
	Group        string      `gorm:"group" json:"group"`                 //分组
	Index        string      `gorm:"index" json:"index"`                 //索引
	ReleaseDate  string      `gorm:"release_date" json:"release_date"`   //发行日期
	Sharpness    string      `gorm:"sharpness" json:"sharpness"`         //清晰度
	Series       string      `gorm:"series" json:"series"`               //系列
	Tags         StringArray `gorm:"tags" json:"tags"`                   //标签
	Length       string      `gorm:"length" json:"length"`               //时长
	Sample       StringArray `gorm:"sample" json:"sample"`               //样板图
	Uncensored   bool        `gorm:"uncensored" json:"uncensored"`       //有码,无码
}

func init() {
	RegisterTable(Info{})
}
