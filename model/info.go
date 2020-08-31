package model

type Info struct {
	BaseModel    `xorm:"extends" json:"-"`
	VideoNo      string      `xorm:"video_no" json:"video_no"`           //编号
	Intro        string      `xorm:"varchar(2048)" json:"intro"`         //简介
	Alias        StringArray `xorm:"json" json:"alias"`                  //别名，片名
	Key          string      `xorm:"key"  json:"key"`                    //秘钥
	M3U8         string      `xorm:"m3u8" json:"m3u8"`                   //M3U8名
	Role         StringArray `xorm:"json" json:"role"`                   //主演
	Director     string      `xorm:"director" json:"director"`           //导演
	Systematics  string      `xorm:"systematics" json:"systematics"`     //分级
	Season       string      `xorm:"season" json:"season"`               //季
	TotalEpisode string      `xorm:"total_episode" json:"total_episode"` //总集数
	Episode      string      `xorm:"episode" json:"episode"`             //集数
	Producer     string      `xorm:"producer" json:"producer"`           //生产商
	Publisher    string      `xorm:"publisher" json:"publisher"`         //发行商
	Type         string      `xorm:"type" json:"type"`                   //类型：film，FanDrama
	Format       string      `xorm:"format" json:"format"`               //输出格式：3D，2D,VR(VR格式：Half-SBS：左右半宽,Half-OU：上下半高,SBS：左右全宽)
	Language     string      `xorm:"language" json:"language"`           //语言
	Caption      string      `xorm:"caption" json:"caption"`             //字幕
	Group        string      `xorm:"group" json:"group"`                 //分组
	Index        string      `xorm:"index" json:"index"`                 //索引
	ReleaseDate  string      `xorm:"release_date" json:"release_date"`   //发行日期
	Sharpness    string      `xorm:"sharpness" json:"sharpness"`         //清晰度
	Series       string      `xorm:"series" json:"series"`               //系列
	Tags         StringArray `xorm:"json tags" json:"tags"`              //标签
	Length       string      `xorm:"length" json:"length"`               //时长
	Sample       StringArray `xorm:"json sample" json:"sample"`          //样板图
	Uncensored   bool        `xorm:"uncensored" json:"uncensored"`       //有码,无码
}

func init() {
	RegisterTable(&Info{})
}
