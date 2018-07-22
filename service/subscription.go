package service

import (
	"time"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rendom/gopodcast/model"
)

type Subscription struct {
	DB *sqlx.DB
}

func (s* Subscription) GetSubscriptionById(user_id int, podcast_id int) (*model.Subscription, error){
	var subscription []model.Subscription

	err := s.DB.Select(&subscription, `SELECT * FROM subscriptions WHERE podcast_id = $1 AND user_id = $2`, podcast_id, user_id)
	if err != nil {
		return nil, err
	}

	if len(subscription) == 0 {
		return nil, errors.New("Failed to find given id")
	}

	return &subscription[0], nil
}

func (s *Subscription) AddSubscription(uid int, podcast_id int) (*model.Subscription, error) {
	var sub = model.Subscription{uid, podcast_id, time.Now()}
	_, err := s.DB.NamedExec(`INSERT INTO subscriptions (podcast_id, user_id, created_at) VALUES(:podcast_id, :user_id, :created_at)`, &sub)
	if err != nil {
		return nil, err
	}

	return &sub, nil
}
