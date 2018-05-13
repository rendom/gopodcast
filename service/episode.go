package service

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rendom/gopodcast/model"
)

type Episode struct {
	DB *sqlx.DB
}

func (e *Episode) New(episode *model.Episode) error {
	_, err := e.DB.NamedExec(
		`INSERT INTO episodes (guid, title, description, url, image, podcast_id, created_at)
		VALUES (:guid, :title, :description, :url, :image, :podcast_id, :created_at)`,
		&episode,
	)

	if err != nil {
		// LOG
		return err
	}

	return nil
}

func (e *Episode) GetPodcastEpisodes(podcastID int) ([]model.Episode, error) {
	episodes := []model.Episode{}
	err := e.DB.Select(
		&episodes,
		"SELECT * FROM episodes WHERE podcast_id = ?",
		podcastID,
	)

	if err != nil {
		return nil, err
	}

	return episodes, nil
}

func (e *Episode) GetEpisodeById(ID int) (*model.Episode, error) {
	var episode model.Episode
	err := e.DB.Get(&episode, "SELECT * FROM episodes WHERE id = ?", ID)

	if err != nil {
		return nil, err
	}

	return &episode, nil
}

func (e *Episode) GetEpisodesByIds(ids []int) ([]model.Episode, error) {
	query, args, err := sqlx.In("SELECT * FROM episodes WHERE id IN (?);", ids)
	if err != nil {
		return nil, err
	}

	var episodes []model.Episode
	query = e.DB.Rebind(query)
	err = e.DB.Select(&episodes, query, args...)

	if err != nil {
		return nil, err
	}

	return episodes, nil
}

func (e *Episode) NewBulk(episodes []model.Episode, podcast_id int) error {
	now := time.Now()
	tx, err := e.DB.Begin()

	for _, v := range episodes {
		tx.Exec(`INSERT OR REPLACE INTO episodes 
		(guid, title, description, url, image, podcast_id, created_at) 
		VALUES (?,?,?,?,?,?,?)`,
			v.GUID, v.Title, v.Description, v.URL, v.ImageURL, podcast_id, now,
		)
	}
	err = tx.Commit()
	return err
}
