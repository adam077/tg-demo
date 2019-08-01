package data

import (
	"github.com/json-iterator/go"
)

type User struct {
	BaseModelUUID
	Name     string `gorm:"column:name;type:text"`
	Password string `gorm:"column:password;type:text"`

	UserScreens []*UserScreen `gorm:"foreignkey:UserId;association_foreignkey:ID"`
}

func (User) TableName() string {
	return "user"
}

func GetUsers(user string) []*User {
	dataConn := GetDataDB("default")
	var temp []*User
	err := dataConn.Find(&temp, "name = ?", user).Error
	if err != nil {
		panic(err)
	}
	return temp
}

func GetUserWithScreens(userId string) []*User {
	dataConn := GetDataDB("default")
	var temp []*User
	err := dataConn.
		Preload("UserScreens").
		Preload("UserScreens.Screen").
		Find(&temp, "id = ?", userId).Error
	if err != nil {
		panic(err)
	}
	return temp
}

type Screen struct {
	BaseModelUUID
	Name    string              `gorm:"column:name;type:text"`
	Content jsoniter.RawMessage `gorm:"column:content;type:json"`
	//Own     int                 `gorm:"column:own;type:int4"` // 代表这个屏是否公有
}

func GetScreens() []*Screen {
	dataConn := GetDataDB("default")
	var temp []*Screen
	err := dataConn.Find(&temp).Error
	if err != nil {
		panic(err)
	}
	return temp
}

func (Screen) TableName() string {
	return "screen"
}

type UserScreen struct {
	BaseModelUUID
	UserId   string `gorm:"column:user_id;type:text"`
	ScreenId string `gorm:"column:screen_id;type:text"` // 用户与屏的归属

	Screen *Screen `gorm:"foreignkey:ScreenId;association_foreignkey:ID"`
}

func (UserScreen) TableName() string {
	return "user_screen"
}

type Component struct {
	BaseModelUUID
	Name    string              `gorm:"column:name;type:text"`
	Path    string              `gorm:"column:path;type:text"`
	Content jsoniter.RawMessage `gorm:"column:content;type:json"`
}

func (Component) TableName() string {
	return "component"
}

func GetComponents() []*Component {
	dataConn := GetDataDB("default")
	var temp []*Component
	err := dataConn.Find(&temp).Error
	if err != nil {
		panic(err)
	}
	return temp
}
