package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-go-go/src/config"
	"sync"
)

var (
	initMutex = &sync.Mutex{}
	dataDBMap = &sync.Map{}
)

func init() {

}

//链接到数据库
func connectDB(dialect, dbname, host string, port int, user, passwd string) (*gorm.DB, error) {
	var connStr string
	var db *gorm.DB
	var err error

	switch dialect {
	case "postgres":
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, passwd, dbname)
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			user, passwd, host, port, dbname)
	default:
		return nil, fmt.Errorf("unknown dialect %s", dialect)
	}

	if db, err = gorm.Open("postgres", connStr); err != nil {
		return nil, err
	}

	return db, nil
}

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
	if db, err := connectDB("postgres", dbName, config.Config.DBConfig.Host, config.Config.DBConfig.Port,
		config.Config.DBConfig.User, config.Config.DBConfig.Password); err != nil {
		panic(err)
	} else {
		db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
		db.DB().SetMaxIdleConns(config.Config.DBConfig.MaxIdleConns)
		db.DB().SetMaxOpenConns(config.Config.DBConfig.MaxOpenConns)
		db.DB().SetConnMaxLifetime(config.Config.DBConfig.MaxLifetime)
		return db, nil
	}
}
