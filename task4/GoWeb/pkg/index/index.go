package index

import "GoStudy/task4/GoWeb/pkg/crawler"

// Index - search index.
// Stores information about documents.

// Interface defines the contract of the search index.
type Interface interface {
	Add(docs []crawler.Document) error
	Search(word string) ([]int, error)
}
