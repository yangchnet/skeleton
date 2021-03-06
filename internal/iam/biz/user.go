package biz

import (
	"context"

	"github.com/yangchnet/skeleton/pkg/logger"
	"github.com/yangchnet/skeleton/tools/passwd"
)

type User struct {
	ID       int64
	Username string
	Password string
	Phone    string
	Email    string
}

// CreateUser hashed password and create a new user.
func (u *IamCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	hashedPasswd, err := passwd.HashPassword(user.Password)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	user.Password = hashedPasswd

	return u.userRepo.CreateUser(ctx, user)
}

func (u *IamCase) GetRolesByUsername(ctx context.Context, username string) ([]*Role, error) {
	return u.userRepo.Roles(ctx, username)
}
