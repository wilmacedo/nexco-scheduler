package main

import (
	"github.com/wilmacedo/nexco-scheduler/configs"
	"github.com/wilmacedo/nexco-scheduler/core"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	runner := core.Runner{}
	runner.Run()
}
