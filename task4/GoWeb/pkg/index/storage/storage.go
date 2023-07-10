// Package storage represents a storage for search robot.
// It allows to store and get information about documents.

package storage

import (
	"GoStudy/task2/GoWeb/pkg/crawler"
	"errors"
	"sort"
	"strings"
)

var (
	NotFound = errors.New("not found")
)

// Index - index where key is a document title word and value is a list of document IDs.
type Index struct {
	data map[string][]int
}

func New() *Index {
	i := Index{
		data: make(map[string][]int),
	}

	return &i
}

// Add adds a document to the index.
func (index *Index) Add(docs []crawler.Document) error {
	for _, doc := range docs {
		for _, word := range strings.Split(doc.Title, " ") {
			word := strings.ToLower(word)
			if !isDuplicate(index.data[word], doc.ID) {
				index.data[word] = append(index.data[word], doc.ID)

				sort.Slice(index.data[word], func(i, j int) bool {
					return index.data[word][i] < index.data[word][j]
				})
			}
		}
	}

	return nil
}

// Search returns a list of document IDs where the word is found.
func (index *Index) Search(word string) ([]int, error) {
	if result := index.data[word]; result == nil {
		return nil, NotFound
	} else {
		return result, nil
	}
}

// isDuplicate checks if the document ID is already in the list.
func isDuplicate(list []int, id int) bool {
	for _, item := range list {
		if item == id {
			return true
		}
	}

	return false
}
