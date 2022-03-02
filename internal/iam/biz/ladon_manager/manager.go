package manager

import (
	"context"
	"encoding/json"

	"github.com/google/wire"
	"github.com/ory/ladon"
	"github.com/yangchnet/skeleton/internal/iam/data/ent"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/authzpolicy"
	"github.com/yangchnet/skeleton/pkg/logger"
)

var ProviderSet = wire.NewSet(NewManager)

var _ ladon.Manager = (*EntManager)(nil)

type EntManager struct {
	client *ent.Client
}

func NewManager(client *ent.Client) ladon.Manager {
	return &EntManager{client: client}
}

// Create persists the policy.
func (l *EntManager) Create(policy ladon.Policy) error {
	panic("not implemented") // TODO: Implement
}

// Update updates an existing policy.
func (l *EntManager) Update(policy ladon.Policy) error {
	panic("not implemented") // TODO: Implement
}

// Get retrieves a policy.
func (l *EntManager) Get(id string) (ladon.Policy, error) {
	panic("not implemented") // TODO: Implement
}

// Delete removes a policy.
func (l *EntManager) Delete(id string) error {
	panic("not implemented") // TODO: Implement
}

// GetAll retrieves all policies.
func (l *EntManager) GetAll(limit int64, offset int64) (ladon.Policies, error) {
	panic("not implemented") // TODO: Implement
}

// FindRequestCandidates returns candidates that could match the request object. It either returns
// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
// the error.
func (l *EntManager) FindRequestCandidates(r *ladon.Request) (ladon.Policies, error) {
	ps, err := l.client.AuthzPolicy.
		Query().
		Where(
			authzpolicy.ObjHasPrefix(r.Resource),
		).All(context.Background())
	if err != nil {
		logger.Error(err)
	}

	res := make([]ladon.Policy, len(ps))
	for i, p := range ps {
		var policy ladon.DefaultPolicy
		if err := json.Unmarshal([]byte(*p.Policy), &policy); err != nil {
			logger.Error(err)
		}
		res[i] = &policy
	}

	return res, nil
}

// FindPoliciesForSubject returns policies that could match the subject. It either returns
// a set of policies that applies to the subject, or a superset of it.
// If an error occurs, it returns nil and the error.
func (l *EntManager) FindPoliciesForSubject(subject string) (ladon.Policies, error) {
	panic("not implemented") // TODO: Implement
}

// FindPoliciesForResource returns policies that could match the resource. It either returns
// a set of policies that apply to the resource, or a superset of it.
// If an error occurs, it returns nil and the error.
func (l *EntManager) FindPoliciesForResource(resource string) (ladon.Policies, error) {
	panic("not implemented") // TODO: Implement
}
