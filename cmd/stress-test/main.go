package main

import (
	"flag"
	"log"

	"github.com/antoniofmoraes/stress-test/internals/services"
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

	res := services.StressTest(*url, *requests, *concurrency)

	logResult(res)
}

func logResult(resultMap map[string]int) {
	total := 0

	log.Println("Amount of responses by status code.")
	for statusCode := range resultMap {
		total += resultMap[statusCode]
		log.Printf("Code %v: %v", statusCode, resultMap[statusCode])
	}

	log.Printf("Total: %v", total)
}
