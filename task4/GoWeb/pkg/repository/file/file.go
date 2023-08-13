// package file implements search repository using a json.

package file

import (
	"GoStudy/task4/GoWeb/pkg/crawler"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const fileRepositoryPath = "/task4/GoWeb/pkg/repository/file/repository.txt"

// Open - opens the repository file for reading and writing.
func Open() (*os.File, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	repoPath := filepath.Join(exPath, fileRepositoryPath)

	file, err := os.OpenFile(repoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Read - reads the repository file.
func Read(reader io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(reader)

	var bytes []byte

	for scanner.Scan() {
		bytes = append(bytes, []byte(scanner.Text())...)
	}

	if error := scanner.Err(); error != nil {
		return nil, error
	}

	return bytes, nil
}

// Write - writes to the repository file.
func Write(writer io.Writer, url string, docs []crawler.Document) error {
	var jsonBytes []byte
	for i, doc := range docs {
		bts, err := json.Marshal(doc)
		if err != nil {
			return err
		}

		bts = append([]byte(fmt.Sprintf(`"%d": `, i)), bts...)
		if i != len(docs)-1 {
			bts = append(bts, []byte(",")...)
		}

		jsonBytes = append(jsonBytes, bts...)
	}
	jsonBytes = []byte(fmt.Sprintf(`{"url": "%s", "docs": {%s}};%v`, url, jsonBytes, "\n"))

	_, err := writer.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}
