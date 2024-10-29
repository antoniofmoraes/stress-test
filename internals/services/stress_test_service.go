package services

import (
	"net/http"
	"strconv"
	"sync"
)

func StressTest(url string, requests int, concurrency int) map[string]int {
	ch := make(chan int, concurrency)
	var responses []string = make([]string, requests)
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < requests; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			responses[i] = sendRequest(url)
		}(i)
	}

	wg.Wait()
	return countResponses(responses)
}

func sendRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "err"
	}
	defer resp.Body.Close()
	return strconv.Itoa(resp.StatusCode)
}

func countResponses(arr []string) map[string]int {
	responseMap := map[string]int{}

	for i := 0; i < len(arr); i++ {
		_, ok := responseMap[arr[i]]
		if !ok {
			responseMap[arr[i]] = 0
		}
		responseMap[arr[i]]++
	}

	return responseMap
}
