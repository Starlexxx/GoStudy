package main

import (
	"GoStudy/task2/GoWeb/pkg/crawler"
	"GoStudy/task2/GoWeb/pkg/crawler/spider"
	"flag"
	"log"
	"strings"
)

var (
	webSites = []string{
		"https://go.dev",
		"https://golang.org",
	}
	searchWord string
)

func init() {
	flag.StringVar(&searchWord, "s", "", "word to search in scanned documents")
	flag.Parse()
}

func main() {
	c := spider.New()
	var docs []crawler.Document

	for _, url := range webSites {
		data, err := c.Scan(url, 2)
		if err != nil {
			log.Printf("GoWeb: %s", err)
			continue
		}
		docs = append(docs, data...)
	}

	for _, d := range docs {
		if searchWord == "" {
			log.Printf("GoWeb: %s - %s", d.URL, d.Title)
		} else if searchWord != "" && strings.Contains(d.URL, searchWord) {
			log.Printf("GoWeb: %s - %s", d.URL, d.Title)
		}
	}
}
