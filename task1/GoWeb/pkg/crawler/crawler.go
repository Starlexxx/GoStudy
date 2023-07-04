package crawler

// Crawler - search robot.
// Performs site scanning.

// Interface defines the contract of the search robot.
type Interface interface {
	Scan(url string, depth int) ([]Document, error)
	BatchScan(urls []string, depth int, workers int) (<-chan Document, <-chan error)
}

// Document - a document, a web page obtained by a search robot.
type Document struct {
	ID    int
	URL   string
	Title string
	Body  string
}
