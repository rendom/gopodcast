package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

func getTime(ti time.Time) *graphql.Time {
	return &graphql.Time{Time: ti}
}
