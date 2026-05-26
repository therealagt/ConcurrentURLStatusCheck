package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	urls := readURLs(os.Args[1:])
	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "usage: go run . <url> [url...]")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	const maxConcurrent = 5
	sem := make(chan struct{}, maxConcurrent)

	for _, raw := range urls {
		url := raw
		wg.Add(1)
		sem <- struct{}{}
		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()
			code, err := checkURL(u)
			if err != nil {
				fmt.Printf("%s ERROR %v\n", u, err)
				return
			}
			fmt.Printf("%s %d\n", u, code)
		}(url)
	}
	wg.Wait()
}

func checkURL(url string) (int, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func readURLs(args []string) []string {
	return args
}
