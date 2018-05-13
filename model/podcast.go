package model

import (
	"time"
)

type Podcast struct {
	ID          int
	GUID        int
	Name        string
	Author      string
	FeedURL     string    `db:"feed_URL"`
	FeedType    string    `db:"feed_type"`
	Description string    `db:"description"`
	ImageURL    string    `db:"image_URL"`
	PubDate     time.Time `db:"pub_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	LatestFetch time.Time `db:"latest_fetch"`
	TTL         int
}
