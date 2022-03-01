package biz

import (
	"context"
	"time"

	"github.com/yangchnet/skeleton/tools/passwd"
)

type User struct {
	ID       int64
	Username string
	Password string
	Phone    string
	Email    string
}

type Token struct {
	Jwt        string
	Expiration time.Duration
}

func (u *IamCase) Token(ctx context.Context, user *User) (*Token, error) {
	panic("not implemented")
}

// CreateUser hashed password and create a new user.
func (u *IamCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	hashedPasswd, err := passwd.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPasswd

	return u.repo.CreateUser(ctx, user)
}