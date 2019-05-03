package config

import (
	"github.com/jinzhu/configor"
	"log"
	"os"
	"path/filepath"
)

var Config = struct {
	PostgresUrl string `env:"POSTGRES_URL" required:"true"`
}{}

var (
	Root = os.Getenv("GOPATH") + "/src/go-go-go"
)

func init() {
	var err error
	if err = configor.Load(&Config, filepath.Join(Root, "etc/config.yaml")); err != nil {
		log.Fatalf("load config error: %v", err)
	}
}
