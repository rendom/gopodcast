package user

type Repostory interface {
	getById(userID int) (User, error)
	update(User) (User, error)

	getSubscriptions(userID int) error
	getInProgressFeedItems(userID int) error
}
