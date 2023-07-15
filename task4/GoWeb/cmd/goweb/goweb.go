package main

import (
	"GoStudy/task4/GoWeb/pkg/crawler"
	"GoStudy/task4/GoWeb/pkg/crawler/spider"
	"GoStudy/task4/GoWeb/pkg/index/storage"
	"errors"
	"flag"
	"log"
	"strings"
)

var (
	webSites = []string{
		"https://go.dev",
		"https://github.com/",
	}
	searchWord string

	UndefinedId = errors.New("undefined id")
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

	s := storage.New()
	err := s.Add(docs)
	if err != nil {
		log.Printf("GoWeb: %s", err)
	}
	if searchWord != "" {
		docIds, err := s.Search(strings.ToLower(searchWord))
		if err != nil {
			log.Printf("GoWeb: %s", err)

			return
		}
		for _, id := range docIds {
			d, err := findDoc(docs, id)
			if err != nil {
				log.Printf("GoWeb: %s", err)
			}
			log.Printf("GoWeb: %s - %s", d.URL, d.Title)
		}
	} else {
		for _, d := range docs {
			log.Printf("GoWeb: %s - %s", d.URL, d.Title)
		}
	}
}

func findDoc(docs []crawler.Document, id int) (crawler.Document, error) {
	low := 0
	high := len(docs) - 1
	for low <= high {
		median := (low + high) / 2
		if docs[median].ID < id {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(docs) || docs[low].ID != id {
		return crawler.Document{}, UndefinedId
	} else {
		return docs[low], nil
	}
}
