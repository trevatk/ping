package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	url := os.Getenv("PING_SERVER_URL")
	if url == "" {
		url = "http://localhost:9090/"
	}

	rq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("failed to create new http request %v", err)
	}

	rq.Header.Add("Content-Type", "application/json")

	rq = rq.WithContext(ctx)

	resp, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Fatalf("failed to send http request %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body %v", err)
	}
	defer resp.Body.Close()

	log.Printf("status code: %d, body: %s", resp.StatusCode, string(body))
}
