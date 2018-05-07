package resolver

import (
	"github.com/rendom/gopodcast/model"
	"github.com/rendom/gopodcast/service"
)

type episodesConnectionResolver struct {
	episodes   []*model.Episode
	totalCount int
	from       *int
	to         *int
}

func (r *episodesConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *episodesConnectionResolver) Edges() *[]*episodesEdgeResolver {
	l := make([]*episodesEdgeResolver, len(r.episodes))
	for i := range l {
		l[i] = &episodesEdgeResolver{
			cursor: service.EncodeCursor(&(r.episodes[i].ID)),
			model:  r.episodes[i],
		}
	}
	return &l
}

func (r *episodesConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: false,
	}
}
