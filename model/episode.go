package model

import "time"

type Episode struct {
	ID          int
	GUID        string
	Title       string
	Description string
	ImageURL    string `db:"image"`
	PodcastID   string `db:"podcast_id"`
	URL         string
	CreatedAt   time.Time `db:"created_at"`
}
