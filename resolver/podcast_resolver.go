package resolver

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/mmcdole/gofeed"
	"github.com/rendom/gopodcast/model"
	"github.com/rendom/gopodcast/service"
)

type podcastResolver struct {
	p              *model.Podcast
	EpisodeService *service.Episode
}

type NewPodcastInput struct {
	URL string
}

func (r *Resolver) AddNewPodcast(ctx context.Context, args NewPodcastInput) (*podcastResolver, error) {
	if ok := ctx.Value(service.ContextAuthIsAuthedKey); ok != true {
		return nil, errors.New("unauthorized")
	}

	pod, err := r.PodcastService.GetPodcastByFeedURL(args.URL)
	if pod != nil {
		return &podcastResolver{pod, r.EpisodeService}, nil
	}

	// var pod *podcast.Podcast
	// decoder := xml.NewDecoder(resp.Body)
	// err = decoder.Decode(pod)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(args.URL)
	if err != nil {
		return nil, errors.New("Invalid rss")
	}

	var title string
	var author string
	var pubdate time.Time
	var description string
	var image string

	title = getTitle(feed)
	author = getAuthor(feed)
	description = getDescription(feed)
	pubdate = *getPublishedDate(feed)
	image = getImage(feed)



/*	m := model.Podcast{
		Name:        feed.Title,
		//Author:      feed.Author.Email, This used to work, need to fix parser or something
		Author:      feed.Author.Email,
		FeedURL:     args.URL,
		PubDate:     *feed.PublishedParsed,
		FeedType:    "rss",
		Description: feed.Description,
		ImageURL:    feed.Image.URL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		LatestFetch: time.Now(),
	}*/


	m := model.Podcast{
		Name:        title,
		Author:      author,
		FeedURL:     args.URL,
		PubDate:     pubdate,
		FeedType:    "rss",
		Description: description,
		ImageURL:    image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		LatestFetch: time.Now(),
	}

	err = r.PodcastService.New(&m)
	if err != nil {
		return nil, err
	}

	pod, err = r.PodcastService.GetPodcastByFeedURL(args.URL)
	if err != nil {
		return nil, err
	}

	var episodes []model.Episode
	for _, v := range feed.Items {
		episodes = append(
			episodes,
			model.Episode{
				GUID:        v.GUID,
				Title:       v.Title,
				Description: v.Description,
				URL:         v.Enclosures[0].URL,
			},
		)
	}
	r.EpisodeService.NewBulk(episodes, pod.ID)

	return &podcastResolver{pod, r.EpisodeService}, nil
}

func (r *Resolver) Podcasts(ctx context.Context) (*[]*podcastResolver, error) {
	if ok := ctx.Value(service.ContextAuthIsAuthedKey); ok != true {
		return nil, errors.New("unauthorized")
	}

	podcasts, err := r.PodcastService.GetPodcasts()
	if err != nil {
		log.Println(err)
		return nil, errors.New("Unable to fetch podcasts")
	}

	var resolvers = make([]*podcastResolver, len(podcasts))
	for k, v := range podcasts {
		resolvers[k] = &podcastResolver{
			&v,
			r.EpisodeService,
		}
	}

	return &resolvers, nil
}

func (r *podcastResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(r.p.ID))
}

func (r *podcastResolver) Name() string {
	return r.p.Name
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

func (r *podcastResolver) Episodes() ([]*episodeResolver, error) {
	episodes, err := r.EpisodeService.GetPodcastEpisodes(r.p.ID)
	if err != nil {
		return nil, err
	}

	var resolvers = make([]*episodeResolver, len(episodes))
	for k, _ := range episodes {
		resolvers[k] = &episodeResolver{&episodes[k]}
	}

	return resolvers, nil
}

// func (r *podcastResolver) TTL() *int {
// 	return &r.p.TTL
// }

func getTitle(feed *gofeed.Feed) (string) {
	title := feed.Title
	if title != "" {
		return title
	}
	return ""
}

func getDescription(feed *gofeed.Feed) (string) {
	return feed.Description
}

func getAuthor(feed *gofeed.Feed) (string) {
	if feed.Author != nil {
		return feed.Author.Email
	} else {
		return ""
	}

}

func getPublishedDate(feed *gofeed.Feed) (*time.Time) {
	if feed.PublishedParsed != nil {
		return feed.PublishedParsed
	} else {
		time := time.Now()
		return &time
	}

}

func getImage(feed *gofeed.Feed) (string) {
	if feed.Image != nil {
		return feed.Image.URL
	} else {
		return ""
	}

}
