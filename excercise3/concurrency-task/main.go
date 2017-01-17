package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"log"
)

var urls = []string{
	"https://www.google.co.uk",
	"http://golang.org",
	"http://www.bbc.co.uk",
	"https://uk.yahoo.com/",
}

func req(ctx context.Context, url string, back chan string) {
	tr := &http.Transport{}
	startTime := time.Now()

	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	duration := time.Since(startTime)
	log.Printf("Time Taken to complete this request: %s ", duration)

	back <- url
}

func makeRequests() {
	ctx, _ := context.WithCancel(context.Background())
	back := make(chan string)
	fmt.Println("Making requests...")
	for _, url := range urls {
		go req(ctx, url, back)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case u := <-back:
			fmt.Printf("%s wins!\n", u)
		}
	}
}

func main() {
	makeRequests()
}
