package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/wilmacedo/nexco-scheduler/configs"

	"github.com/go-co-op/gocron"
)

func RunJobs() {
	scheduler := gocron.NewScheduler(time.UTC)
	env := configs.GetEnv()

	scheduler.Every(env.Interval).Minutes().Do(func() {
		urls := strings.Split(env.Urls, ",")

		for _, url := range urls {
			err := fetcher(url)
			if err != nil {
				fmt.Printf("[%s] url=%s err=%v\n", time.Now().Format("2006-01-02 15:04:05"), url, err)
			}
		}
	})

	scheduler.StartBlocking()
}
