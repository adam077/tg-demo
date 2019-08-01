package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModelUUID struct {
	ID string `gorm:"column:id;type:char(36);primary_key;not null"`
	BaseModel
}

type BaseModelIncrementID struct {
	ID int64 `gorm:"column:id;type:bigserial;primary_key;not null"`
	BaseModel
}

type BaseModel struct {
	CreatedAt time.Time `gorm:"column:create_time;type:timestamp with time zone" json:"-"`
	UpdatedAt time.Time `gorm:"column:update_time;type:timestamp with time zone" json:"-"`
}

func UpdateOne(dbConn *gorm.DB, target interface{}, targetMap map[string]interface{}) error { // with id
	if targetMap == nil || len(targetMap) == 0 {
		return nil
	}
	return dbConn.Model(target).Update(targetMap).Error
}

func AddOne(dbConn *gorm.DB, target interface{}) error {
	return dbConn.Create(target).Error
}

func DeleteOne(dbConn *gorm.DB, target interface{}) error {
	return dbConn.Delete(target).Error
}
