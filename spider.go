package main

import (
	"errors"
	"flag"
	"fmt"
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

func main() {
	url, err := getInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("Url to crawl is: %s\n", url)
	}
}
