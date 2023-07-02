// Package spider represents a web site content scanner.
// It allows to get a list of links and page titles inside the web site by its URL.
package spider

import (
	"GoStudy/task2/GoWeb/pkg/crawler"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Service - search robot service.
type Service struct{}

// New - constructor of search robot service.
func New() *Service {
	s := Service{}
	return &s
}

// Scan recursively traverse linked of the website specified in URL
// taking into account the depth of transition by links, passed in depth.
func (s *Service) Scan(url string, depth int) (data []crawler.Document, err error) {
	pages := make(map[string]string)

	err = parse(url, url, depth, pages)
	if err != nil {
		return nil, err
	}

	for url, title := range pages {
		item := crawler.Document{
			URL:   url,
			Title: title,
		}
		data = append(data, item)
	}

	return data, nil
}

// parse recursively traverses the links on the page passed in url.
// The depth of recursion is set in depth.
// Each found link is written to the associative array
// data along with the page name.
func parse(url, baseurl string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	data[url] = pageTitle(page)

	if depth == 1 {
		return nil
	}
	links := pageLinks(nil, page)
	for _, link := range links {
		link = strings.TrimSuffix(link, "/")
		// relative link
		if strings.HasPrefix(link, "/") && len(link) > 1 {
			link = baseurl + link
		}
		// link is already scanned
		if data[link] != "" {
			continue
		}
		// link includes baseurl
		if strings.HasPrefix(link, baseurl) {
			err := parse(link, baseurl, depth-1, data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// pageTitle recursively traverses the HTML page and returns the value of the <title> element.
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLinks recursively scans the HTML page nodes and returns all found links without duplicates.
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// sliceContains returns true if the array contains the passed value
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
