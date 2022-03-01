package biz

import "context"

type IamRepo interface {
	// ============================================
	// User
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type IamCase struct {
	repo IamRepo
}

func NewIamCase(ctx context.Context, repo IamRepo) *IamCase {
	return &IamCase{repo: repo}
}
