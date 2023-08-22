package config

var availableJobs = []string{"fetcher"}

func GetJobs() []string {
	return availableJobs
}
