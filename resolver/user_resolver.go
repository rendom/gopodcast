package resolver

import (
	"errors"

	"github.com/rendom/gopodcast/model"
)

type UserInput struct {
	Name     string
	Email    string
	Password string
}

func (r *Resolver) CreateUser(args UserInput) (string, error) {
	h, err := model.PasswordHash(args.Password)
	if err != nil {
		return "", err
	}

	tmpu := model.User{
		Name:     args.Name,
		Email:    args.Email,
		Password: h,
	}
	err = r.UserService.CreateUser(tmpu)

	if err != nil {
		return "", err
	}

	u, err := r.UserService.GetUserByEmail(tmpu.Email)
	if err != nil {
		return "", err
	}

	t, err := r.AuthService.CreateToken(u.ID)
	if err != nil {
		return "", errors.New(
			`Account created but unable to create login token,
			Try to login with your new account`,
		)
	}

	return t, nil
}

func (r *Resolver) Login(args UserInput) (string, error) {
	u, err := r.UserService.GetUserByEmail(args.Email)
	invalidErr := errors.New("Invalid")
	if err != nil {
		return "", invalidErr
	}

	err = u.CheckPassword(args.Password)
	if err != nil {
		return "", invalidErr
	}

	t, err := r.AuthService.CreateToken(u.ID)
	if err != nil {
		// Log this
		return "", errors.New("Unable to create login token, try again later")
	}

	return t, nil
}
