package model

import "time"

type Episode struct {
	ID          int
	Title       string
	Description string
	ImageURL    string
	URL         string
	CreatedAt   time.Time
}
