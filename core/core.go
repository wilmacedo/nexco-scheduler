package core

import (
	"fmt"
	"time"

	"github.com/wilmacedo/nexco-scheduler/configs"

	"github.com/go-co-op/gocron"
)

type Runner struct{}

func (runner Runner) Run() {
	scheduler := gocron.NewScheduler(time.UTC)
	env := configs.GetEnv()

	scheduler.Every(env.Interval).Second().Do(func() {
		fmt.Printf("Interval: %d\n", env.Interval)
	})

	scheduler.StartBlocking()
}
