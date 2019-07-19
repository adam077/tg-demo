package config

import (
	"fmt"
	"os"
)

var Config = Cf{
	PostgresUrl: "postgres://postgres:password@139.180.202.66:5432/%s?sslmode=disable",
}

func init() {
	FillEnvWithString("POSTGRES_URL", &Config.PostgresUrl, false)
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
