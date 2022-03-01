package data

import (
	"context"

	"github.com/yangchnet/skeleton/internal/echo/data/ent"
	"github.com/yangchnet/skeleton/internal/iam/biz"
	"github.com/yangchnet/skeleton/pkg/logger"
)

const (
	StatusActive  = "active"
	StatusDisable = "disable"
	StatusDelete  = "delete"
)

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
