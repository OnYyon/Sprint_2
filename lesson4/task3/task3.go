package task3

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int
	Err        error
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	responses := make([]*APIResponse, len(urls))
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(url string, index int) {
			resp := &APIResponse{URL: url}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				resp.StatusCode = 0
				resp.Data = ""
				resp.Err = err
				responses[index] = resp
			}
			response, err := http.DefaultClient.Do(req)
			if err != nil {
				resp.StatusCode = response.StatusCode
				resp.Data = ""
				if ctx.Err() == context.DeadlineExceeded {
					resp.Err = context.DeadlineExceeded
				} else {
					resp.Err = err
				}
				responses[index] = resp
			}
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				resp.StatusCode = response.StatusCode
				resp.Data = ""
				resp.Err = err
				responses[index] = resp
				return
			}
			resp.StatusCode = response.StatusCode
			resp.Data = string(body)
			resp.Err = nil
			responses[index] = resp
		}(urls[i], i)
	}
	wg.Wait()
	return responses
}
