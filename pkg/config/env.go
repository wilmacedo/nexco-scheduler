package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	Interval int
}

var env *Environment
var defaultEnv = &Environment{
	Interval: 5,
}

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return ErrFailedToLoadEnv
	}

	envInterval := os.Getenv("INTERVAL")
	var interval int
	if len(envInterval) == 0 {
		interval = defaultEnv.Interval
	} else {
		convertedInterval, err := strconv.Atoi(envInterval)
		if err != nil {
			return ErrConvertInterval
		}

		interval = convertedInterval
	}

	environment := &Environment{
		Interval: interval,
	}

	env = environment

	return nil
}

func GetEnv() *Environment {
	return env
}
