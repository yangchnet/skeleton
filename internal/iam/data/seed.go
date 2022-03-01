package data

import (
	"context"

	"github.com/yangchnet/skeleton/internal/iam/data/ent"
	"github.com/yangchnet/skeleton/pkg/logger"
)

func seed(ctx context.Context, client *ent.Client) error {
	return seedRoot(ctx, client)
}

func seedRoot(ctx context.Context, client *ent.Client) error {
	root := struct {
		Username string
		Passwd   string
		Phone    string
		Email    string
		Status   string
	}{
		Username: "root",
		Passwd:   "$2a$10$KPoGEhY51u31lY8AI1Q27eBLmDr6ImrUPJWqik5fPPP6RxlBrsaI6",
		Phone:    "17354277557",
		Email:    "lich.mailer@gmail.com",
		Status:   StatusActive,
	}

	_, err := client.User.Create().
		SetUsername(root.Username).
		SetPasswd(root.Passwd).
		SetPhone(root.Phone).
		SetEmail(root.Email).
		SetStatus(root.Status).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil
		}
		logger.Error("seed root user failed ", err)

		return err
	}

	return nil
}
