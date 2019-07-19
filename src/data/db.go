package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-go-go/src/config"
	"sync"
	"time"
)

var (
	initMutex = &sync.Mutex{}
	dataDBMap = &sync.Map{}
)

func GetDataDB(dbName string) *gorm.DB {
	initMutex.Lock()
	defer func() {
		initMutex.Unlock()
	}()
	if v, ok := dataDBMap.Load(dbName); !ok {
		if db, err := initDataDB(dbName); err != nil {
			panic(err)
		} else {
			dataDBMap.Store(dbName, db)

			return db
		}
	} else {
		db := v.(*gorm.DB)
		return db
	}
}

func initDataDB(dbName string) (*gorm.DB, error) {
	conn := fmt.Sprintf(config.Config.PostgresUrl, dbName)
	if db, err := gorm.Open("postgres", conn); err != nil {
		return nil, err
	} else {
		db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
		db.DB().SetMaxIdleConns(16)
		db.DB().SetMaxOpenConns(32)
		maxLifetime, _ := time.ParseDuration("30m")
		db.DB().SetConnMaxLifetime(maxLifetime)
		return db, nil
	}
}
