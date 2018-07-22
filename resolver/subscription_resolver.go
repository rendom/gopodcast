package resolver

import (
	"context"
	"errors"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rendom/gopodcast/model"
	"github.com/rendom/gopodcast/service"
)

type  subscriptionResolver struct {
	s	*model.Subscription
}

type SubscribeInput struct {
	ID int32
}

func (r *Resolver) Subscriptions(ctx context.Context) (*[]*podcastResolver, error) {
	if ok := ctx.Value(service.ContextAuthIsAuthedKey); ok != true {
		return nil, errors.New("unauthorized")
	}

	return nil, errors.New("not implemented")
}

func (r *Resolver) Subscribe(ctx context.Context, args SubscribeInput) (*subscriptionResolver, error) {
	if ok := ctx.Value(service.ContextAuthIsAuthedKey); ok != true {
		return nil, errors.New("Unauthorized")
	}
	uId := ctx.Value(service.ContextAuthUIDKey).(int)
	podId := int(args.ID)


	pod, err := r.PodcastService.GetPodcastById(podId)
	if pod == nil {
		return nil, errors.New("Failed to find podcast with the given id")
	}

	sub, err := r.SubscriptionService.GetSubscriptionById(uId, podId)
	if sub != nil {
		return &subscriptionResolver{sub}, nil
	}

	sub, err = r.SubscriptionService.AddSubscription(ctx.Value(service.ContextAuthUIDKey).(int), int(args.ID))
	if err != nil {
		return nil, err
	}
	return &subscriptionResolver{sub}, nil
}

func (r *subscriptionResolver) Podcastid() graphql.ID {
	return graphql.ID(strconv.Itoa(r.s.PodcastID))
}
