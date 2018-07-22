package model

import (
	"time"
)

type Subscription struct {
	UID         int `db:"user_id"`
	PodcastID   int `db:"podcast_id"`
	CreatedAt   time.Time `db:"created_at"`
}
