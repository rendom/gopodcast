package resolver

import (
	"context"
	"errors"

	"github.com/rendom/gopodcast/service"
)

func (r *Resolver) Subscriptions(ctx context.Context) (*[]*podcastResolver, error) {
	if ok := ctx.Value(service.ContextAuthIsAuthedKey); ok != true {
		return nil, errors.New("unauthorized")
	}

	return nil, errors.New("not implemented")
}
