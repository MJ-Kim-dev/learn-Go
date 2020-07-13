package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFail = errors.New("Request Faild")

type result struct {
	url   string
	check bool
}

func main() {

	c := make(chan result)
	urls := []string{
		"https://google.com",
		"https://naver.com",
		"https://instagram.com",
		"https://facebook.com",
		"https://amazon.com",
	}

	for _, url := range urls {
		fmt.Println("Checking: ", url)
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		c <- result{url: url, check: false}
	}
	c <- result{url: url, check: true}
}
