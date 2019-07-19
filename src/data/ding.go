package data

import "github.com/rs/zerolog/log"

type Ding struct {
	Name   string `gorm:"column:name;type:text"`
	ChatId string `gorm:"column:chat_id;type:text"`
}

func (Ding) TableName() string {
	return "ding"
}

func GetDingChatId(name string) []*Ding {
	var bizDb = GetDataDB("default")
	var temp []*Ding
	if bizDb.Find(&temp, "name = ?", name).Error != nil {
		log.Info().Msg("数据访问失败")
	}
	return temp
}
