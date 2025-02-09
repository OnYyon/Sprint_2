package task4

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://127.0.0.1:8081/provideData")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(response.StatusCode)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}

func StartServer(maxTimeout time.Duration) {
	mux := http.NewServeMux()
	timeoutHandler := http.TimeoutHandler(http.HandlerFunc(ResponseHandler), maxTimeout, "StatusServiceUnavailable")
	mux.Handle("/readSource", timeoutHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
