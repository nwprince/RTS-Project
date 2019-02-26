package downloader

type Media struct {
	Prepared  bool   `json:"prepared"`
	MediaType string `json:"mediaType"`
	Title     string `json:"title"`
	Container string `json:"container"`
	ID        string `json:"id"`
}
