package main

import (
	"net/http"
	"time"
)

type Result struct {
	success     bool
	statusCode  int
	latencyMs   float64
	errorString string
}

func makeReq(method, url string) Result {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return Result{success: false, errorString: err.Error()}
	}

	start := time.Now()
	resp, err := client.Do(req)
	latency := time.Since(start).Seconds() * 1000 // in ms

	if err != nil {
		return Result{success: false, latencyMs: latency, errorString: err.Error()}
	}

	defer resp.Body.Close()
	return Result{
		success:    resp.StatusCode >= 200 && resp.StatusCode < 300,
		statusCode: resp.StatusCode,
		latencyMs:  latency,
	}
}
