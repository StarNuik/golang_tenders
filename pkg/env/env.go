package env

import (
	"os"
)

type env struct {
	PostgresUrl string
	ServerUrl   string
}

func Read() env {
	out := env{
		ServerUrl: "0.0.0.0:8080",
	}
	if url, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		out.ServerUrl = url
	}
	if url, ok := os.LookupEnv("POSTGRES_CONN"); ok {
		out.PostgresUrl = url
	}
	return out
}
