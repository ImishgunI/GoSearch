package main

import (
	"GoSearch/pkg/crawler"
	"GoSearch/pkg/crawler/index"
	"GoSearch/pkg/crawler/spider"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func Scanner(url1, url2 string) (slice []crawler.Document) {
	crawle := spider.Service{}
	firstUrl, err := crawle.Scan(url1, 2)
	if err != nil {
		log.Fatalf("cannot scan this url %s", url1)
	}

	slice = append(slice, firstUrl...)

	secondUrl, err := crawle.Scan(url2, 2)
	if err != nil {
		log.Fatalf("cannot scan this url %s", url2)
	}

	slice = append(slice, secondUrl...)

	return slice
}

func printUrls(urls []crawler.Document, flagname string) {
	for _, url := range urls {
		if strings.Contains(url.URL, flagname) {
			fmt.Println(url.URL)
		}
	}
}

func indexing(urls []crawler.Document) (ri []index.ReverseIndex, err error) {
	var idx int
	for _, url := range urls {
		item := index.ReverseIndex{
			ID:    idx,
			URL:   url.URL,
			Title: url.Title,
		}
		idx++
		ri = append(ri, item)
	}
	return ri, nil
}

func main() {
	url1 := "https://go.dev"
	url2 := "https://golang.org"

	var flagVar string
	flag.StringVar(&flagVar, "s", "", "for searching substring in urls")

	flag.Parse()

	if flagVar == "" {
		log.Fatalf("flag needs an argument -s, but u have %s", flagVar)
		os.Exit(2)
	}
	urls := Scanner(url1, url2)

	ri, err := indexing(urls)
	if err != nil {
		log.Fatalln("reverse index weren't completed")
	}
	_ = ri

	printUrls(urls, flagVar)
}
