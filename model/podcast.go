package model

import (
	"time"
)

type Podcast struct {
	ID          int
	Name        string
	Author      string
	FeedURL     string
	FeedType    string
	Description string
	ImageURL    string
	PubDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LatestFetch time.Time
	TTL         int
}
