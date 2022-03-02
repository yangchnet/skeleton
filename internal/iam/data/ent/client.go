// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/yangchnet/skeleton/internal/iam/data/ent/migrate"

	"github.com/yangchnet/skeleton/internal/iam/data/ent/authzpolicy"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/role"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/tenant"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AuthzPolicy is the client for interacting with the AuthzPolicy builders.
	AuthzPolicy *AuthzPolicyClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Tenant is the client for interacting with the Tenant builders.
	Tenant *TenantClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AuthzPolicy = NewAuthzPolicyClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.Tenant = NewTenantClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		AuthzPolicy: NewAuthzPolicyClient(cfg),
		Role:        NewRoleClient(cfg),
		Tenant:      NewTenantClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		AuthzPolicy: NewAuthzPolicyClient(cfg),
		Role:        NewRoleClient(cfg),
		Tenant:      NewTenantClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AuthzPolicy.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AuthzPolicy.Use(hooks...)
	c.Role.Use(hooks...)
	c.Tenant.Use(hooks...)
	c.User.Use(hooks...)
}

// AuthzPolicyClient is a client for the AuthzPolicy schema.
type AuthzPolicyClient struct {
	config
}

// NewAuthzPolicyClient returns a client for the AuthzPolicy from the given config.
func NewAuthzPolicyClient(c config) *AuthzPolicyClient {
	return &AuthzPolicyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `authzpolicy.Hooks(f(g(h())))`.
func (c *AuthzPolicyClient) Use(hooks ...Hook) {
	c.hooks.AuthzPolicy = append(c.hooks.AuthzPolicy, hooks...)
}

// Create returns a create builder for AuthzPolicy.
func (c *AuthzPolicyClient) Create() *AuthzPolicyCreate {
	mutation := newAuthzPolicyMutation(c.config, OpCreate)
	return &AuthzPolicyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AuthzPolicy entities.
func (c *AuthzPolicyClient) CreateBulk(builders ...*AuthzPolicyCreate) *AuthzPolicyCreateBulk {
	return &AuthzPolicyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AuthzPolicy.
func (c *AuthzPolicyClient) Update() *AuthzPolicyUpdate {
	mutation := newAuthzPolicyMutation(c.config, OpUpdate)
	return &AuthzPolicyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthzPolicyClient) UpdateOne(ap *AuthzPolicy) *AuthzPolicyUpdateOne {
	mutation := newAuthzPolicyMutation(c.config, OpUpdateOne, withAuthzPolicy(ap))
	return &AuthzPolicyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthzPolicyClient) UpdateOneID(id int) *AuthzPolicyUpdateOne {
	mutation := newAuthzPolicyMutation(c.config, OpUpdateOne, withAuthzPolicyID(id))
	return &AuthzPolicyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AuthzPolicy.
func (c *AuthzPolicyClient) Delete() *AuthzPolicyDelete {
	mutation := newAuthzPolicyMutation(c.config, OpDelete)
	return &AuthzPolicyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AuthzPolicyClient) DeleteOne(ap *AuthzPolicy) *AuthzPolicyDeleteOne {
	return c.DeleteOneID(ap.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AuthzPolicyClient) DeleteOneID(id int) *AuthzPolicyDeleteOne {
	builder := c.Delete().Where(authzpolicy.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthzPolicyDeleteOne{builder}
}

// Query returns a query builder for AuthzPolicy.
func (c *AuthzPolicyClient) Query() *AuthzPolicyQuery {
	return &AuthzPolicyQuery{
		config: c.config,
	}
}

// Get returns a AuthzPolicy entity by its id.
func (c *AuthzPolicyClient) Get(ctx context.Context, id int) (*AuthzPolicy, error) {
	return c.Query().Where(authzpolicy.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthzPolicyClient) GetX(ctx context.Context, id int) *AuthzPolicy {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCreateBy queries the create_by edge of a AuthzPolicy.
func (c *AuthzPolicyClient) QueryCreateBy(ap *AuthzPolicy) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ap.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(authzpolicy.Table, authzpolicy.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authzpolicy.CreateByTable, authzpolicy.CreateByColumn),
		)
		fromV = sqlgraph.Neighbors(ap.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AuthzPolicyClient) Hooks() []Hook {
	return c.hooks.AuthzPolicy
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Create returns a create builder for Role.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Role.
func (c *RoleClient) QueryUsers(r *Role) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, role.UsersTable, role.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// TenantClient is a client for the Tenant schema.
type TenantClient struct {
	config
}

// NewTenantClient returns a client for the Tenant from the given config.
func NewTenantClient(c config) *TenantClient {
	return &TenantClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tenant.Hooks(f(g(h())))`.
func (c *TenantClient) Use(hooks ...Hook) {
	c.hooks.Tenant = append(c.hooks.Tenant, hooks...)
}

// Create returns a create builder for Tenant.
func (c *TenantClient) Create() *TenantCreate {
	mutation := newTenantMutation(c.config, OpCreate)
	return &TenantCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tenant entities.
func (c *TenantClient) CreateBulk(builders ...*TenantCreate) *TenantCreateBulk {
	return &TenantCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tenant.
func (c *TenantClient) Update() *TenantUpdate {
	mutation := newTenantMutation(c.config, OpUpdate)
	return &TenantUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TenantClient) UpdateOne(t *Tenant) *TenantUpdateOne {
	mutation := newTenantMutation(c.config, OpUpdateOne, withTenant(t))
	return &TenantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TenantClient) UpdateOneID(id int) *TenantUpdateOne {
	mutation := newTenantMutation(c.config, OpUpdateOne, withTenantID(id))
	return &TenantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tenant.
func (c *TenantClient) Delete() *TenantDelete {
	mutation := newTenantMutation(c.config, OpDelete)
	return &TenantDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TenantClient) DeleteOne(t *Tenant) *TenantDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TenantClient) DeleteOneID(id int) *TenantDeleteOne {
	builder := c.Delete().Where(tenant.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TenantDeleteOne{builder}
}

// Query returns a query builder for Tenant.
func (c *TenantClient) Query() *TenantQuery {
	return &TenantQuery{
		config: c.config,
	}
}

// Get returns a Tenant entity by its id.
func (c *TenantClient) Get(ctx context.Context, id int) (*Tenant, error) {
	return c.Query().Where(tenant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TenantClient) GetX(ctx context.Context, id int) *Tenant {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Tenant.
func (c *TenantClient) QueryUsers(t *Tenant) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tenant.Table, tenant.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tenant.UsersTable, tenant.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TenantClient) Hooks() []Hook {
	return c.hooks.Tenant
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPolicys queries the policys edge of a User.
func (c *UserClient) QueryPolicys(u *User) *AuthzPolicyQuery {
	query := &AuthzPolicyQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(authzpolicy.Table, authzpolicy.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PolicysTable, user.PolicysColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoles queries the roles edge of a User.
func (c *UserClient) QueryRoles(u *User) *RoleQuery {
	query := &RoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.RolesTable, user.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBelong queries the belong edge of a User.
func (c *UserClient) QueryBelong(u *User) *TenantQuery {
	query := &TenantQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.BelongTable, user.BelongColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
