package main

import (
	"flag"
)

type input struct {
	Url string
	Max int
}

func (inp *input) Parse() {
	url := flag.String("url", "http://www.google.com", "the url of the site to be crawled")
	max := flag.Int("max", 1, "max number of requests to make")

	flag.Parse()

	inp.Url = *url
	inp.Max = *max
}

func (inp input) Success() bool {
	return inp.Url != "" && inp.Max > 0
}
