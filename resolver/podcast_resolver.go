package resolver

import (
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rendom/gopodcast/model"
)

type podcastResolver struct {
	p *model.Podcast
}

var lol []*podcastResolver

func (r *Resolver) Podcasts() (*[]*podcastResolver, error) {
	var resolvers = make([]*podcastResolver, 5)
	for k, _ := range resolvers {
		resolvers[k] = &podcastResolver{
			p: &model.Podcast{
				ID:      k,
				Name:    fmt.Sprintf("Podcast %d", k),
				Author:  fmt.Sprintf("Skaning %d", k),
				FeedURL: fmt.Sprintf("https://skaning.com/v%d/rss", k),
			},
		}
	}

	return &resolvers, nil
}

func (r *podcastResolver) ID() *graphql.ID {
	i := graphql.ID(r.p.ID)
	return &i
}

func (r *podcastResolver) Name() *string {
	return &r.p.Name
}

func (r *podcastResolver) Author() *string {
	return &r.p.Author
}

func (r *podcastResolver) FeedURL() *string {
	return &r.p.FeedURL
}

func (r *podcastResolver) FeedType() *string {
	return &r.p.FeedType
}

func (r *podcastResolver) Description() *string {
	return &r.p.Description
}

func (r *podcastResolver) ImageURL() *string {
	return &r.p.ImageURL
}

func (r *podcastResolver) PubDate() *graphql.Time {
	return getTime(r.p.PubDate)
}

func (r *podcastResolver) CreatedAt() *graphql.Time {
	return getTime(r.p.CreatedAt)
}

func (r *podcastResolver) UpdatedAt() *graphql.Time {
	return getTime(r.p.UpdatedAt)
}

func (r *podcastResolver) LatestFetch() *graphql.Time {
	return getTime(r.p.UpdatedAt)
}

// func (r *podcastResolver) TTL() *int {
// 	return &r.p.TTL
// }
