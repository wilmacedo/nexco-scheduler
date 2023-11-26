package core

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/wilmacedo/nexco-scheduler/configs"

	"github.com/go-co-op/gocron"
)

type Runner struct{}

func (runner Runner) Run() {
	scheduler := gocron.NewScheduler(time.UTC)
	env := configs.GetEnv()

	scheduler.Every(env.Interval).Second().Do(func() {
		output, err := exec.Command("pwd").Output()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Interval: %d\n", env.Interval)
		fmt.Printf("Output: %s", output)
	})

	scheduler.StartBlocking()
}
