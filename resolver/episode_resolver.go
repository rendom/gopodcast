package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rendom/gopodcast/model"
)

type episodeResolver struct {
	e *model.Episode
}

func (r *episodeResolver) ID() graphql.ID {
	return graphql.ID(r.e.ID)
}

func (r *episodeResolver) Title() string {
	return r.e.Title
}

func (r *episodeResolver) Description() *string {
	return &r.e.Description
}

func (r *episodeResolver) URL() *string {
	return &r.e.URL
}

func (r *episodeResolver) Image() *string {
	return &r.e.ImageURL
}

func (r *episodeResolver) CreatedAt() *graphql.Time {
	return getTime(r.e.CreatedAt)
}
