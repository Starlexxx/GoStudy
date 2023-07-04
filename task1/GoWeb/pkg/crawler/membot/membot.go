package membot

import "GoStudy/task1/GoWeb/pkg/crawler"

// Service - imitation of a search robot service.
type Service struct{}

// New - constructor of imitation of a search robot service.
func New() *Service {
	s := Service{}
	return &s
}

// Scan - returns a pre-prepared data set
func (s *Service) Scan(url string, depth int) ([]crawler.Document, error) {

	data := []crawler.Document{
		{
			ID:    0,
			URL:   "https://yandex.ru",
			Title: "Яндекс",
		},
		{
			ID:    1,
			URL:   "https://google.ru",
			Title: "Google",
		},
	}

	return data, nil
}
