package configs

import (
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	Interval int
}

var env *Environment

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	mapEnv, err := godotenv.Read(".env")
	if err != nil {
		return nil
	}

	interval := mapEnv["INTERVAL"]
	if interval == "" {
		interval = "5"
	}

	convertedInterval, err := strconv.Atoi(interval)
	if err != nil {
		return err
	}

	env = &Environment{
		Interval: convertedInterval,
	}

	return nil
}

func GetEnv() *Environment {
	return env
}
