package main

import (
	"go-cron/pkg/config"
	"go-cron/pkg/core"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	runner := core.Runner{}
	runner.Run()
}
