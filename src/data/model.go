package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableTest struct {
	BaseModelUUID
	Test string `gorm:"column:test;type:text"`
}

func (TableTest) TableName() string {
	return "table_test"
}

type BaseModelUUID struct {
	ID string `gorm:"column:id;type:char(32);primary_key;not null"`
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

func (*BaseModel) Indexes() map[string][]string {
	return map[string][]string{}
}

func (*BaseModel) UniqueIndexes() map[string][]string {
	return map[string][]string{}
}

func MigrateTable(db *gorm.DB, tableModel ModelInterface) {
	db.AutoMigrate(tableModel)

	for k, v := range tableModel.Indexes() {
		db.Model(tableModel).AddIndex(k, v...)
	}

	for k, v := range tableModel.UniqueIndexes() {
		db.Model(tableModel).AddUniqueIndex(k, v...)
	}
}

type ModelInterface interface {
	Indexes() map[string][]string
	UniqueIndexes() map[string][]string
}
