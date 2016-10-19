package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
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

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting %s; aborting!\n", url)
		return
	}
	fmt.Printf("URL retrieved: %s\n", url)
	uChan <- url

	body := resp.Body
	defer body.Close()
	parser := html.NewTokenizer(body)

	for {
		token := parser.Next()

		switch {
		case token == html.ErrorToken:
			// EOF
			return
		case token == html.StartTagToken:
			t := parser.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						if attr.Val != "#" && attr.Val != "javascript:void(0)" {
							uChan <- attr.Val
						}
					}
				}
			}
		}
	}
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

	max := 100
	doneCount := 0
	started := 1
	for doneCount < started {
		select {
		case <-doneChan:
			doneCount++
		case link := <-urlChan:
			fmt.Printf(": %s\n", link)
			foundUrls = append(foundUrls, link)
			if started < max {
				started++
				go getLinks(link, urlChan, doneChan)
			}
		default:
			continue
		}
	}

	if len(foundUrls) > 0 {
		fmt.Printf("Found %d links!\n", len(foundUrls))
	}
}
