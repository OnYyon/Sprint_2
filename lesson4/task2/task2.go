package task2

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string
	StatusCode int
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &APIResponse{Data: string(body), StatusCode: response.StatusCode}, nil
}
