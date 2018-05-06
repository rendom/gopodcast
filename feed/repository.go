package feed

type Repository interface {
	getFeedByUrl(string) (Feed, error)
	addFeed(Feed) (Feed, error)
	updateFeed(Feed) (Feed, error)
}
