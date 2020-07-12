package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFail = errors.New("Request Faild")

func main() {

	var results = make(map[string]bool)

	urls := []string{
		"https://google.com",
		"https://naver.com",
		"https://instagram.com",
		"https://facebook.com",
		"https://amazon.com",
	}

	for _, url := range urls {
		fmt.Println("Checking: ", url)
		err := hitURL(url)
		result := true

		if err != nil {
			fmt.Println(err)
			result = false
		}
		results[url] = result

	}
	for url, result := range results {
		fmt.Println(url, result)
	}

}

func hitURL(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFail
	}
	return nil
}
