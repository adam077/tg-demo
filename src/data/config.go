package data

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
	"time"
)

var Env = Cf{
	PostgresUrl: "postgres://postgres:password@139.180.202.66:5432/%s?sslmode=disable",
}

const (
	// for config
	Scheduler = "scheduler"
)

var configMap sync.Map

func GetConfig(config string) string {
	if result, ok := configMap.Load(config); ok {
		return result.(string)
	}
	return ""

}

func init() {
	FillEnvWithString("POSTGRES_URL", &Env.PostgresUrl, false)
	UpdateConfig()
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				UpdateConfig()
			}
		}
	}()
}

func FillEnvWithString(env string, value *string, required bool) {
	if envValue, exist := os.LookupEnv("POSTGRES_URL"); exist {
		*value = envValue
	} else if required {
		panic(fmt.Sprintf("no env: %s", env))
	}
}

type Cf struct {
	PostgresUrl string
}

func UpdateConfig() {
	configs := GetConfigs()
	for x := range configs {
		configMap.Store(configs[x].Name, configs[x].Content)
	}
}

type ConfigTable struct {
	Name    string `gorm:"column:name;type:text"`
	Content string `gorm:"column:content;type:text"`
	Example string `gorm:"column:example;type:text"`
}

func (ConfigTable) TableName() string {
	return "config"
}

func GetConfigs() []*ConfigTable {
	var bizDb = GetDataDB("config")
	var temp []*ConfigTable
	if bizDb.Find(&temp).Error != nil {
		log.Info().Msg("数据访问失败")
	}
	return temp
}
