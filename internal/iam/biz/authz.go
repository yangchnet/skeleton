package biz

import (
	"context"

	"github.com/ory/ladon"
)

func (uc *IamCase) CanDo(ctx context.Context, username string, resource string, action string) bool {
	warden := &ladon.Ladon{Manager: uc.policyManager}
	roles, err := uc.userRepo.Roles(ctx, username)
	if err != nil {
		return false
	}

	for _, role := range roles {
		ladonReq := &ladon.Request{
			Resource: resource,
			Action:   action,
			Subject:  role.RoleName,
		}

		if err := warden.IsAllowed(ladonReq); err == nil {
			return true
		}
	}

	return false
}
