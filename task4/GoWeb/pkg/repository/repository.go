package repository

import (
	"GoStudy/task4/GoWeb/pkg/crawler"
	"GoStudy/task4/GoWeb/pkg/repository/file"
	"bytes"
	"encoding/json"
)

// SearchResult - returns an array of documents by url
func SearchResult(url string) ([]crawler.Document, error) {
	repo, err := file.Open()
	if err != nil {
		return nil, err
	}

	bts, err := file.Read(repo)
	if err != nil {
		return nil, err
	}

	var docs []crawler.Document

	for _, jsonStr := range bytes.Split(bts, []byte(";")) {
		if string(jsonStr) == "" {
			continue
		}
		docsMap := make(map[string]interface{})
		err := json.Unmarshal(jsonStr, &docsMap)
		if err != nil {
			return nil, err
		}

		if docsMap["url"] == url {
			for _, doc := range docsMap["docs"].(map[string]interface{}) {
				docStr, err := json.Marshal(doc)
				if err != nil {
					return nil, err
				}

				doc := crawler.Document{}
				err = json.Unmarshal(docStr, &doc)
				if err != nil {
					return nil, err
				}

				docs = append(docs, doc)
			}
		}
	}

	return docs, nil
}

// Add - adds a document to the repository
func Add(url string, docs []crawler.Document) error {
	if isDuplicatedURL(url) {
		return nil
	}

	repo, err := file.Open()
	if err != nil {
		return err
	}
	defer repo.Close()

	err = file.Write(repo, url, docs)
	if err != nil {
		return err
	}

	return nil
}

// isDuplicatedURL - checks if the url is already in the repository
func isDuplicatedURL(url string) bool {
	repo, err := file.Open()
	if err != nil {
		return false
	}
	defer repo.Close()

	bts, err := file.Read(repo)
	if err != nil {
		return false
	}

	for _, str := range bytes.Split(bts, []byte(";")) {
		var data map[string]interface{}

		err := json.Unmarshal(str, &data)
		if err != nil {
			return false
		}

		if data["url"] == url {
			return true
		}
	}

	return false
}
