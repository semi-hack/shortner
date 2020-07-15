package models

const (
	// Linkcollection ...
	Linkcollection = "Links"
)

// Link ...
type Link struct {
	Short string `json: "short"`
	URL string `json: "url"`
}