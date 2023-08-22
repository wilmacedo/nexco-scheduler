package core

import (
	"fmt"
	"os/exec"
	"time"

	"go-cron/pkg/config"

	"github.com/go-co-op/gocron"
)

type Runner struct{}

func (runner Runner) Run() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(config.GetEnv().Interval).Second().Do(func() {
		output, err := exec.Command("pwd").Output()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Output: %s", output)
	})

	scheduler.StartBlocking()
}
