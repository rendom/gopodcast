package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// PasswordHash generate hash with bcrypt.
func PasswordHash(pw string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(h), nil
}

// CheckPassword compare User.password hash with inputed password
func (u *User) CheckPassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
}
