package data

import "github.com/rs/zerolog/log"

type EatWhatTable struct {
	Name string `gorm:"column:name;type:text"`
	Num  int    `gorm:"column:num;type:int8"`
}

func (EatWhatTable) TableName() string {
	return "eat_what"
}

func GetEatNames() []*EatWhatTable {
	var bizDb = GetDataDB("default")
	var temp []*EatWhatTable
	if bizDb.Find(&temp).Order("num").Error != nil {
		log.Info().Msg("数据访问失败")
	}
	return temp
}
