package service

import (
	"encoding/base64"
	"fmt"
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
)

func EncodeCursor(i *int) graphql.ID {
	return graphql.ID(
		base64.StdEncoding.EncodeToString(
			[]byte(fmt.Sprintf("cursor%d", *i)),
		),
	)
}

func DecodeCursor(after *string) (*string, error) {
	var decodedValue string
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		i := strings.TrimPrefix(string(b), "cursor")
		decodedValue = i
	}
	return &decodedValue, nil
}
