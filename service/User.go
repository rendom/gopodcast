package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rendom/gopodcast/model"
)

// type User interface {
// 	getUser(int) *model.User
// 	createUser(name string, email string, pw string) (user *model.User)
// 	login(email string, pw string) (token string)
// }

type User struct {
	DB *sqlx.DB
}

func (u *User) GetUserById(ID string) (*model.User, error) {
	return u.getUserByCol("id", ID)
}

func (u *User) GetUserByEmail(email string) (*model.User, error) {
	return u.getUserByCol("email", email)
}

func (u *User) CreateUser(user model.User) error {
	_, err := u.DB.NamedExec(
		`INSERT INTO users (name, email, password)
		VALUES (:name, :email, :password)`,
		&user,
	)

	return err
}

func (u *User) getUserByCol(col string, v string) (*model.User, error) {
	var user = model.User{}
	query := fmt.Sprintf(
		`SELECT id, email, name, password 
		FROM users
		WHERE %s = $1`,
		col,
	)

	err := u.DB.Get(&user, query, v)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
