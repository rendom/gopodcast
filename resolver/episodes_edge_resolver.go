package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rendom/gopodcast/model"
)

type episodesEdgeResolver struct {
	cursor graphql.ID
	model  *model.Episode
}

func (r *episodesEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *episodesEdgeResolver) Node() *episodeResolver {
	return &episodeResolver{e: r.model}
}
