package config

import (
	"os"
	"strconv"
)

type Env struct {
	PORT int16

	DB_HOST string
	DB_PORT int16
	DB_USER string
	DB_PASS string
	DB_NAME string

	RUN_AUTO_MIGRATION bool
}

func ReadEnv() *Env {
	env := &Env{
		PORT:     int16(getEnvAsInt("PORT")),

		DB_HOST:  os.Getenv("DB_HOST"),
		DB_PORT:  int16(getEnvAsInt("DB_PORT")),
		DB_USER:  os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_NAME:  os.Getenv("DB_NAME"),
		RUN_AUTO_MIGRATION: getEnvAsBool("RUN_AUTO_MIGRATION"),
	}

	return env
}


func getEnvAsInt(key string) int {
	value, _ := strconv.Atoi(os.Getenv(key))
	return value
}

func getEnvAsBool(key string) bool {
	value, _ := strconv.ParseBool(os.Getenv(key))
	return value
}
