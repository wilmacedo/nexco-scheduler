package core

import (
	"fmt"
	"net/http"
	"strings"
)

func fetcher(url string) error {
	url = strings.Trim(url, " ")

	data, err := http.Get(url)
	if err != nil {
		return err
	}
	defer data.Body.Close()

	if data.StatusCode != http.StatusOK {
		return fmt.Errorf("access rejected from %s", url)
	}

	return nil
}
