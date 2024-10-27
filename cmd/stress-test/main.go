package main

import (
	"flag"
	"log"
)

func main() {
	url := flag.String("url", "", "URL of the service tested. Required.")
	requests := flag.Int("requests", 1000, "Total number of requests. Default: 1000")
	concurrency := flag.Int("concurrency", 10, "Number of simultaneous requests. Default: 10")
	flag.Parse()

	if *url == "" {
		log.Fatal("The URL flag is required. Example: [run command] -url=https://google.com")
	}
	if *concurrency > *requests {
		log.Fatal("The concurrency number needs to be less than the number of requests.")
	}

	log.Printf("URL: %s. Requests: %v. Concurrency: %v.", *url, *requests, *concurrency)
}
