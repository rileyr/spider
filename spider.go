package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	url := flag.String("url", "", "the url of the site to be crawled")

	flag.Parse()

	if *url == "" {
		fmt.Printf("Url is required.\n")
		os.Exit(1)
	} else {
		fmt.Printf("Url to be crawled: %v\n", *url)
	}
}
