package config

import (
	"github.com/jinzhu/configor"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Config = struct {
	ServicePort int `env:"SERVICE_PORT" default:"8080"`

	Debug           bool `env:"DEBUG" default:"false"`
	EnableCORS      bool `env:"ENABLE_CORS" default:"false"`
	EnableScheduler bool `env:"ENABLE_SCHEDULER" default:"false"`

	DBConfig struct {
		Host     string `env:"BUSINESS_DB_HOST"`
		Port     int    `env:"BUSINESS_DB_PORT"`
		User     string `env:"BUSINESS_DB_USER" required:"true"`
		Password string `env:"BUSINESS_DB_PASSWORD" required:"true"`

		MaxIdleConns int `env:"BUSINESS_DB_MAX_IDLE_CONNS" required:"true"`
		MaxOpenConns int `env:"BUSINESS_DB_MAX_OPEN_CONNS" required:"true"`

		MaxLifetimeStr string `env:"BUSINESS_DB_MAX_LIFETIME" default:"30m"`
		MaxLifetime    time.Duration
	}
}{}

var (
	Root = os.Getenv("GOPATH") + "/src/go-go-go"
)

func init() {
	var err error
	if err = configor.Load(&Config, filepath.Join(Root, "etc/config.yaml")); err != nil {
		log.Fatalf("load config error: %v", err)
	}
	if Config.DBConfig.MaxLifetime, err = time.ParseDuration(Config.DBConfig.MaxLifetimeStr); err != nil {
		log.Fatalf("parse Config.BusinessDB.MaxLifetime error: %s", err.Error())
	}
}
