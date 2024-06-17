package main

import (
	"GoSearch/pkg/crawler"
	"GoSearch/pkg/crawler/spider"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	Функция для поиска по сайтам go.dev и golang.org
	нужно передать входные аргуметны командной строки
	создать флаг -s, просканить на соответствие флагу
	и дальше запустить скан по сайтам и вывести все
	ссылки на запрос.
*/

func Scanner(url1, url2 string) (slice []crawler.Document) {
	crawle := spider.Service{}
	firstUrl, err := crawle.Scan(url1, 4)
	if err != nil {
		log.Fatalf("cannot scan this url %s", url1)
	}

	slice = append(slice, firstUrl...)

	secondUrl, err := crawle.Scan(url2, 4)
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

	printUrls(urls, flagVar)
}
