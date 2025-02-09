package task1

import (
	"net/http"
	"time"
)

func FetchURL(url string) string {
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	response, err := http.Get(url)
	if err != nil {
		return "Failed to fetch"
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return "Successfully fetched"
	} else {
		return "Failed to fetch"
	}
}
