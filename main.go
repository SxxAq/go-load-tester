package main

import (
	"flag"
	"fmt"

)



func main() {
	totalRequests := flag.Int("n", 100, "Total number of requests")
	concurrency := flag.Int("c", 100, "Number of concurrent workers")
	method := flag.String("m", "GET", "HTTP method")
	url := flag.String("url", "", "Target URL")

	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required. Use --url to provide one.")
		return
	}

	fmt.Println("Starting Load Test")
	fmt.Printf(
		"Requests: %d | Concurrency: %d | Method: %s | URL: %s\n",
		*totalRequests,
		*concurrency,
		*method,
		*url,
	)
}
