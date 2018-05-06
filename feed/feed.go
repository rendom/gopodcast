package feed

import "time"

type Feed struct {
	ID   int
	URL  string // Unique identifier?
	Type string // rss(podcast)..

	Name        string
	Description string
	Image       string

	CreatedAt time.Time
	UpdatedAt time.Time
	FetchedAt time.Time
}
