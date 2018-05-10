package resolver

import "github.com/rendom/gopodcast/service"

type Resolver struct {
	UserService *service.User
	AuthService *service.AuthService
}
