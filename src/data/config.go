package data

import (
	"fmt"
	"os"
)

var Env = Cf{
	PostgresUrl: "postgres://postgres:password@localhost:5432/%s?sslmode=disable",
}

func init() {
	FillEnv("POSTGRES_URL", &Env.PostgresUrl, false)
}

func FillEnv(env string, value *string, required bool) {
	if envValue, exist := os.LookupEnv(env); exist {
		*value = envValue
	} else if required {
		panic(fmt.Sprintf("no env: %s", env))
	}
}

type Cf struct {
	PostgresUrl string
}
