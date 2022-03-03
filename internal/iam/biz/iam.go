package biz

import (
	"context"

	"github.com/ory/ladon"
	"github.com/yangchnet/skeleton/pkg/token"
)

type IamRepo interface {
	// ============================================
	// User
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	Roles(ctx context.Context, username string) ([]*Role, error)
	Passwd(ctx context.Context, username string) (string, error)
}

type RoleRepo interface{}

type BindRepo interface{}

type TenantRepo interface{}

type PolicyRepo interface{}

type IamCase struct {
	userRepo      UserRepo
	roleRepo      RoleRepo
	bindRepo      BindRepo
	tenantRepo    TenantRepo
	policyRepo    PolicyRepo
	policyManager ladon.Manager
	tokenMaker    token.Maker
}

func NewIamCase(
	ctx context.Context,
	userRepo UserRepo,
	roleRepo RoleRepo,
	bindRepo BindRepo,
	tenantRepo TenantRepo,
	policyRepo PolicyRepo,
	manager ladon.Manager,
	tokenMaker token.Maker) *IamCase {
	return &IamCase{
		userRepo:      userRepo,
		roleRepo:      roleRepo,
		bindRepo:      bindRepo,
		tenantRepo:    tenantRepo,
		policyRepo:    policyRepo,
		policyManager: manager,
		tokenMaker:    tokenMaker,
	}
}
