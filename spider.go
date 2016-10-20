package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func stringInCollection(c []string, s string) bool {
	for _, e := range c {
		if s == e {
			return true
		}
	}
	return false
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

	body := resp.Body
	defer body.Close()
	parser := html.NewTokenizer(body)

	skips := []string{"", "#", "/", "javascript:void(0)"}

	for {
		token := parser.Next()

		switch {
		case token == html.ErrorToken:
			return
		case token == html.StartTagToken:
			t := parser.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						if url != attr.Val && !stringInCollection(skips, attr.Val) {
							uChan <- attr.Val
						}
					}
				}
			}
		}
	}
}

func main() {
	inp := new(input)
	inp.Parse()
	if !inp.Success() {
		fmt.Println("fail")
		os.Exit(1)
	} else {
		fmt.Printf("URL to crawl is: %s\n", inp.Url)
	}

	foundUrls := make([]string, 0, 50)

	urlChan := make(chan string)

	doneChan := make(chan bool)

	go getLinks(inp.Url, urlChan, doneChan)

	doneCount := 0
	started := 1
	for doneCount < started {
		select {
		case <-doneChan:
			doneCount++
		case link := <-urlChan:
			fmt.Printf(": %s\n", link)
			foundUrls = append(foundUrls, link)
			if started < inp.Max {
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
