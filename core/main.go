package core

import (
	"fmt"
	"net/http"
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
			url = strings.Trim(url, " ")

			data, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer data.Body.Close()

			if data.StatusCode != http.StatusOK {
				fmt.Printf("[%s] url=%s code=%d\n", time.Now().Format("2006-01-02 15:04:05"), url, data.StatusCode)
			}
		}
	})

	scheduler.StartBlocking()
}
