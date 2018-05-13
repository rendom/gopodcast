package resolver

import "github.com/rendom/gopodcast/service"

type Resolver struct {
	UserService    *service.User
	PodcastService *service.Podcast
	EpisodeService *service.Episode
	AuthService    *service.AuthService
}
