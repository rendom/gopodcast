package service

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/rendom/gopodcast/model"
)

type Podcast struct {
	DB *sqlx.DB
}

func (p *Podcast) New(podcast *model.Podcast) error {
	_, err := p.DB.NamedExec(
		`INSERT INTO podcasts (name, description, author, feed_URL, pub_date, created_at, updated_at, image_URL, latest_fetch)
		VALUES (:name, :description, :author, :feed_URL, :pub_date, :created_at, :updated_at, :image_URL, :latest_fetch)`,
		&podcast,
	)

	if err != nil {
		// LOG
		return err
	}

	return nil
}

func (p *Podcast) GetPodcasts() ([]model.Podcast, error) {
	var podcasts []model.Podcast

	err := p.DB.Select(&podcasts, "SELECT * FROM podcasts")
	if err != nil {
		return nil, err
	}

	return podcasts, nil
}

func (p *Podcast) getPodcastByCol(col string, v string) (*model.Podcast, error) {
	var podcast = model.Podcast{}
	query := fmt.Sprintf(
		`SELECT *
		FROM podcasts
		WHERE %s = $1`,
		col,
	)

	err := p.DB.Get(&podcast, query, v)
	if err != nil {
		return nil, err
	}

	return &podcast, nil
}

func (p *Podcast) GetPodcastById(ID int) (*model.Podcast, error) {
	return p.getPodcastByCol("id", strconv.Itoa(ID))
}

func (p *Podcast) GetPodcastByFeedURL(feed string) (*model.Podcast, error) {
	stmt, err := p.DB.Preparex("SELECT * FROM podcasts WHERE feed_URL = ?")

	if err != nil {
		return nil, err
	}

	var pod model.Podcast
	err = stmt.Get(&pod, feed)
	if err != nil {
		return nil, err
	}

	return &pod, nil
}

func (p *Podcast) GetPodcastsByIds(ids []int) ([]model.Podcast, error) {
	query, args, err := sqlx.In("SELECT * FROM podcasts WHERE id IN (?);", ids)
	if err != nil {
		return nil, err
	}

	var podcasts []model.Podcast
	query = p.DB.Rebind(query)
	err = p.DB.Select(&podcasts, query, args...)

	if err != nil {
		return nil, err
	}

	return podcasts, nil
}

func (p *Podcast) UserSubscriptions(UID int) ([]model.Podcast, error) {
	var ids []int

	err := p.DB.Select(&ids, "SELECT * FROM subscriptions WHERE user_id = ?", UID)
	if err != nil {
		return nil, err
	}

	return p.GetPodcastsByIds(ids)
}
