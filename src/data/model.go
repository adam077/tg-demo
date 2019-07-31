package data

import (
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

func UpdateOne(target interface{}, targetMap map[string]interface{}) error { // with id
	dbConn := GetDataDB("default")
	if targetMap == nil || len(targetMap) == 0 {
		return nil
	}
	return dbConn.Model(target).Update(targetMap).Error
}

func AddOne(target interface{}) error {
	dbConn := GetDataDB("default")
	return dbConn.Create(target).Error
}

func DeleteOne(target interface{}) error {
	dbConn := GetDataDB("default")
	err := dbConn.Delete(target).Error
	if err != nil {
		return err
	}
	return nil
}
