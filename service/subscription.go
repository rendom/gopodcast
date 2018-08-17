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

func (s *Subscription) GetAllUserSubscriptions(user_id int) ([]model.Subscription, error) {
	var subscriptions []model.Subscription

	err := s.DB.Select(&subscriptions, `SELECT * FROM subscriptions WHERE user_id = $1`,  user_id)
	if err != nil {
		return nil, err
	}

	if len(subscriptions) == 0 {
		return nil, errors.New("Failed to find user subscriptions")
	}
	return subscriptions, nil
}

func (s *Subscription) GetSubscriptionById(user_id int, podcast_id int) (*model.Subscription, error){
	var subscription []model.Subscription

	err := s.DB.Select(&subscription, `SELECT * FROM subscriptions WHERE podcast_id = $1 AND user_id = $2`, podcast_id, user_id)
	if err != nil {
		return nil, err
	}

	if len(subscription) == 0 {
		return nil, errors.New("Failed to find subscription by the given ids")
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


