package data

import (
	"context"

	"github.com/yangchnet/skeleton/internal/echo/data/ent"
	"github.com/yangchnet/skeleton/internal/iam/biz"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/user"
	"github.com/yangchnet/skeleton/pkg/logger"
)

// CreateUser creates a new user.
func (r *Data) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	u, err := r.db.User.
		Create().
		SetUsername(user.Username).
		SetPasswd(user.Password).
		SetPhone(user.Phone).
		SetEmail(user.Email).
		SetStatus(StatusDisable).
		Save(ctx)
	if err != nil {
		logger.Error("Error creating user: ", err)
		if ent.IsConstraintError(err) {
			return nil, biz.ErrExistUsername
		}

		return nil, err
	}
	user.ID = int64(u.ID)

	return user, nil
}

func (r *Data) Roles(ctx context.Context, username string) ([]*biz.Role, error) {
	user, err := r.db.User.Query().Where(
		user.Username(username),
	).Only(ctx)
	if err != nil {
		logger.Error("failed found username.", err)
		if ent.IsNotFound(err) {
			return nil, biz.ErrNotFound
		}

		return nil, err
	}
	roles, err := user.QueryRoles().All(ctx)
	if err != nil {
		logger.Error("failed found corresponding roles.", err)
		if ent.IsNotFound(err) {
			return nil, biz.ErrNotFound
		}

		return nil, err
	}

	res := make([]*biz.Role, 0, len(roles))
	for _, role := range roles {
		res = append(res, &biz.Role{
			RoleName: role.Rolename,
			Status:   role.Status,
		})
	}

	return res, nil
}
