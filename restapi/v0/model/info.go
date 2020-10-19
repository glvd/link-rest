package model

import "github.com/glvd/link-rest/restapi/common/model"

type Info struct {
	model.BaseModel `json:"-" swaggerignore:"true"`
	Title           string            `gorm:"column:title" json:"title"`                 //标题
	VideoNo         string            `gorm:"column:video_no" json:"video_no"`           //编号
	Intro           string            `gorm:"column:intro;size:2048" json:"intro"`       //简介
	Alias           model.StringArray `gorm:"column:alias" json:"alias"`                 //别名，片名
	Role            model.StringArray `gorm:"column:role" json:"role"`                   //主演
	Director        string            `gorm:"column:director" json:"director"`           //导演
	Systematics     string            `gorm:"column:systematics" json:"systematics"`     //分级
	Season          string            `gorm:"column:season" json:"season"`               //季
	TotalEpisode    string            `gorm:"column:total_episode" json:"total_episode"` //总集数
	Episode         string            `gorm:"column:episode" json:"episode"`             //集数
	Producer        string            `gorm:"column:producer" json:"producer"`           //生产商
	Publisher       string            `gorm:"column:publisher" json:"publisher"`         //发行商
	MediaType       string            `gorm:"column:media_type" json:"media_type"`       //类型：Film,FanDrama,TVDrama,18X
	Format          string            `gorm:"column:format" json:"format"`               //输出格式：3D，2D,VR(VR格式：Half-SBS：左右半宽,Half-OU：上下半高,SBS：左右全宽)
	Language        string            `gorm:"column:language" json:"language"`           //语言
	Caption         string            `gorm:"column:caption" json:"caption"`             //字幕
	Group           string            `gorm:"column:group" json:"group"`                 //分组
	Index           string            `gorm:"column:index" json:"index"`                 //索引
	ReleaseDate     string            `gorm:"column:release_date" json:"release_date"`   //发行日期
	Sharpness       string            `gorm:"column:sharpness" json:"sharpness"`         //清晰度
	Series          string            `gorm:"column:series" json:"series"`               //系列
	Tags            model.StringArray `gorm:"column:tags" json:"tags"`                   //标签
	Length          string            `gorm:"column:length" json:"length"`               //时长
	Sample          model.StringArray `gorm:"column:sample" json:"sample"`               //样板图
	Uncensored      bool              `gorm:"column:uncensored" json:"uncensored"`       //有码,无码
}

func init() {
	model.RegisterTable(Info{})
}
