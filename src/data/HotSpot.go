package data

import "time"

type WeiBoHotSpotMinuteReport struct {
	ID int64 `gorm:"column:id;type:bigserial;primary_key;not null"`

	LogDate string `gorm:"column:log_date;type:char(10)"`
	LogHour int    `gorm:"column:log_hour;type:int4"`
	LogMin  int    `gorm:"column:log_min;type:int4"`

	Content string `gorm:"column:content;type:text"`
	Rank    int    `gorm:"column:rank;type:int4"`

	RequestTS time.Time `gorm:"column:request_ts;type:timestamp with time zone"`
	BaseModel
}

func (WeiBoHotSpotMinuteReport) TableName() string {
	return "weibo_hot_spot_minute_report"
}

type BaiDuHotSpotMinuteReport struct {
	ID int64 `gorm:"column:id;type:bigserial;primary_key;not null"`

	LogDate string `gorm:"column:log_date;type:char(10)"`
	LogHour int    `gorm:"column:log_hour;type:int4"`
	LogMin  int    `gorm:"column:log_min;type:int4"`

	Content  string `gorm:"column:content;type:text"`
	Rank     int    `gorm:"column:rank;type:int4"`
	Searches int    `gorm:"column:searches;type:int4"`

	RequestTS time.Time `gorm:"column:request_ts;type:timestamp with time zone"`
	BaseModel
}

func (BaiDuHotSpotMinuteReport) TableName() string {
	return "baidu_hot_spot_minute_report"
}

type ZhiHuHotSpotMinuteReport struct {
	ID int64 `gorm:"column:id;type:bigserial;primary_key;not null"`

	LogDate string `gorm:"column:log_date;type:char(10)"`
	LogHour int    `gorm:"column:log_hour;type:int4"`
	LogMin  int    `gorm:"column:log_min;type:int4"`

	Content     string  `gorm:"column:content;type:text"`
	Rank        int     `gorm:"column:rank;type:int4"`
	Score       float64 `gorm:"column:score;type:float8"`
	AnswerCount int     `gorm:"column:answer;type:int4"`
	Link        string  `gorm:"column:link;type:text"`

	RequestTS time.Time `gorm:"column:request_ts;type:timestamp with time zone"`
	BaseModel
}

func (ZhiHuHotSpotMinuteReport) TableName() string {
	return "zhihu_hot_spot_minute_report"
}
