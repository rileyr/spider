package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	//	"golang.org/x/net/html"
)

func getInput() (string, error) {
	url := flag.String("url", "", "the url of the site to be crawled")

	flag.Parse()

	if *url == "" {
		return "", errors.New("Missing URL to crawl.")
	}
	return *url, nil
}

func getLinks(url string, uChan chan string, dChan chan bool) {
	defer func() {
		dChan <- true
	}()

	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting %s; aborting!\n", url)
		return
	}
	fmt.Printf("URL retrieved: %s\n", url)
	uChan <- url
}

func main() {
	// Get the URL to start crawling:
	url, err := getInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("URL to crawl is: %s\n", url)
	}

	// collection of urls found on the page:
	foundUrls := make([]string, 0, 50)

	// channel to send urls down:
	urlChan := make(chan string)

	// channel to indicate that we've finished with a page:
	doneChan := make(chan bool)

	// get the links from the input page:
	go getLinks(url, urlChan, doneChan)
	doneCount := 0
	for doneCount < 1 {
		select {
		case <-doneChan:
			doneCount++
		case link := <-urlChan:
			foundUrls = append(foundUrls, link)
		default:
			continue
		}
	}

	if len(foundUrls) > 0 {
		fmt.Printf("Found %d links!\n", len(foundUrls))
	}
}
