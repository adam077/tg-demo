package data

import (
	"testing"
)

func TestMigrateTable(t *testing.T) {
	t1()
}

func t1() {
	db := GetDataDB("default")
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Screen{})
	db.AutoMigrate(&Component{})
	db.AutoMigrate(&UserScreen{})
	db.AutoMigrate(&UserWidget{})
}
