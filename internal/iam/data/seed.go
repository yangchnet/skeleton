package data

import (
	"context"
	"encoding/json"

	"github.com/ory/ladon"
	"github.com/yangchnet/skeleton/internal/iam/data/ent"
	"github.com/yangchnet/skeleton/pkg/logger"
)

type seedHelper struct {
	fns []func(ctx context.Context, client *ent.Client) error
	err error
}

func newSeedHelper(fns ...func(ctx context.Context, client *ent.Client) error) *seedHelper {
	return &seedHelper{
		fns: fns,
		err: nil,
	}
}

func (h *seedHelper) run(ctx context.Context, client *ent.Client) error {
	for _, fn := range h.fns {
		if h.err != nil {
			logger.Fatal("failed to seed, service break. ", h.err)

			return h.err
		}
		h.err = fn(ctx, client)
	}

	return nil
}

func seed(ctx context.Context, client *ent.Client) error {
	// add here seed func
	h := newSeedHelper(seedRootUser, seedPolicy)

	return h.run(ctx, client)
}

func seedRootUser(ctx context.Context, client *ent.Client) error {
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

	user, err := client.User.Create().
		SetUsername(root.Username).
		SetPasswd(root.Passwd).
		SetPhone(root.Phone).
		SetEmail(root.Email).
		SetStatus(root.Status).
		Save(ctx)
	if err != nil {
		return err
	}

	rootRole := struct {
		RoleName string
		Status   string
	}{
		RoleName: "root",
		Status:   StatusActive,
	}

	_, err = client.Role.Create().
		SetRolename(rootRole.RoleName).
		SetStatus(rootRole.Status).
		AddUsers(user).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func seedPolicy(ctx context.Context, client *ent.Client) error {
	policy := ladon.DefaultPolicy{
		Description: "root policy",
		Subjects:    []string{"root"},
		Effect:      ladon.AllowAccess,
		Resources:   []string{"system<.*>"},
		Actions:     []string{"delete", "create", "update", "get"},
		Meta:        []byte{},
	}
	policyJson, err := json.Marshal(policy)
	if err != nil {
		return err
	}

	rootPolicy := struct {
		PolicyName string
		Obj        string
		Policy     string
		Status     string
	}{
		PolicyName: "root policy",
		Obj:        "system::*",
		Policy:     string(policyJson),
		Status:     StatusActive,
	}

	_, err = client.AuthzPolicy.
		Create().
		SetPolicyName(rootPolicy.PolicyName).
		SetObj(rootPolicy.Obj).
		SetPolicy(rootPolicy.Policy).
		SetStatus(rootPolicy.Status).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
