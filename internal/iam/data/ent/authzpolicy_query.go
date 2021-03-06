// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/authzpolicy"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/predicate"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/user"
)

// AuthzPolicyQuery is the builder for querying AuthzPolicy entities.
type AuthzPolicyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthzPolicy
	// eager-loading edges.
	withCreateBy *UserQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthzPolicyQuery builder.
func (apq *AuthzPolicyQuery) Where(ps ...predicate.AuthzPolicy) *AuthzPolicyQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit adds a limit step to the query.
func (apq *AuthzPolicyQuery) Limit(limit int) *AuthzPolicyQuery {
	apq.limit = &limit
	return apq
}

// Offset adds an offset step to the query.
func (apq *AuthzPolicyQuery) Offset(offset int) *AuthzPolicyQuery {
	apq.offset = &offset
	return apq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (apq *AuthzPolicyQuery) Unique(unique bool) *AuthzPolicyQuery {
	apq.unique = &unique
	return apq
}

// Order adds an order step to the query.
func (apq *AuthzPolicyQuery) Order(o ...OrderFunc) *AuthzPolicyQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// QueryCreateBy chains the current query on the "create_by" edge.
func (apq *AuthzPolicyQuery) QueryCreateBy() *UserQuery {
	query := &UserQuery{config: apq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(authzpolicy.Table, authzpolicy.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authzpolicy.CreateByTable, authzpolicy.CreateByColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AuthzPolicy entity from the query.
// Returns a *NotFoundError when no AuthzPolicy was found.
func (apq *AuthzPolicyQuery) First(ctx context.Context) (*AuthzPolicy, error) {
	nodes, err := apq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authzpolicy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *AuthzPolicyQuery) FirstX(ctx context.Context) *AuthzPolicy {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthzPolicy ID from the query.
// Returns a *NotFoundError when no AuthzPolicy ID was found.
func (apq *AuthzPolicyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authzpolicy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *AuthzPolicyQuery) FirstIDX(ctx context.Context) int {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthzPolicy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AuthzPolicy entity is not found.
// Returns a *NotFoundError when no AuthzPolicy entities are found.
func (apq *AuthzPolicyQuery) Only(ctx context.Context) (*AuthzPolicy, error) {
	nodes, err := apq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authzpolicy.Label}
	default:
		return nil, &NotSingularError{authzpolicy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *AuthzPolicyQuery) OnlyX(ctx context.Context) *AuthzPolicy {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthzPolicy ID in the query.
// Returns a *NotSingularError when exactly one AuthzPolicy ID is not found.
// Returns a *NotFoundError when no entities are found.
func (apq *AuthzPolicyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = &NotSingularError{authzpolicy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *AuthzPolicyQuery) OnlyIDX(ctx context.Context) int {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthzPolicies.
func (apq *AuthzPolicyQuery) All(ctx context.Context) ([]*AuthzPolicy, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return apq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (apq *AuthzPolicyQuery) AllX(ctx context.Context) []*AuthzPolicy {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthzPolicy IDs.
func (apq *AuthzPolicyQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := apq.Select(authzpolicy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *AuthzPolicyQuery) IDsX(ctx context.Context) []int {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *AuthzPolicyQuery) Count(ctx context.Context) (int, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return apq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (apq *AuthzPolicyQuery) CountX(ctx context.Context) int {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *AuthzPolicyQuery) Exist(ctx context.Context) (bool, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return apq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *AuthzPolicyQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthzPolicyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *AuthzPolicyQuery) Clone() *AuthzPolicyQuery {
	if apq == nil {
		return nil
	}
	return &AuthzPolicyQuery{
		config:       apq.config,
		limit:        apq.limit,
		offset:       apq.offset,
		order:        append([]OrderFunc{}, apq.order...),
		predicates:   append([]predicate.AuthzPolicy{}, apq.predicates...),
		withCreateBy: apq.withCreateBy.Clone(),
		// clone intermediate query.
		sql:  apq.sql.Clone(),
		path: apq.path,
	}
}

// WithCreateBy tells the query-builder to eager-load the nodes that are connected to
// the "create_by" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AuthzPolicyQuery) WithCreateBy(opts ...func(*UserQuery)) *AuthzPolicyQuery {
	query := &UserQuery{config: apq.config}
	for _, opt := range opts {
		opt(query)
	}
	apq.withCreateBy = query
	return apq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PolicyName string `json:"policy_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthzPolicy.Query().
//		GroupBy(authzpolicy.FieldPolicyName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (apq *AuthzPolicyQuery) GroupBy(field string, fields ...string) *AuthzPolicyGroupBy {
	group := &AuthzPolicyGroupBy{config: apq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return apq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PolicyName string `json:"policy_name,omitempty"`
//	}
//
//	client.AuthzPolicy.Query().
//		Select(authzpolicy.FieldPolicyName).
//		Scan(ctx, &v)
//
func (apq *AuthzPolicyQuery) Select(fields ...string) *AuthzPolicySelect {
	apq.fields = append(apq.fields, fields...)
	return &AuthzPolicySelect{AuthzPolicyQuery: apq}
}

func (apq *AuthzPolicyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range apq.fields {
		if !authzpolicy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	return nil
}

func (apq *AuthzPolicyQuery) sqlAll(ctx context.Context) ([]*AuthzPolicy, error) {
	var (
		nodes       = []*AuthzPolicy{}
		withFKs     = apq.withFKs
		_spec       = apq.querySpec()
		loadedTypes = [1]bool{
			apq.withCreateBy != nil,
		}
	)
	if apq.withCreateBy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, authzpolicy.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AuthzPolicy{config: apq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := apq.withCreateBy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*AuthzPolicy)
		for i := range nodes {
			if nodes[i].user_policys == nil {
				continue
			}
			fk := *nodes[i].user_policys
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_policys" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CreateBy = n
			}
		}
	}

	return nodes, nil
}

func (apq *AuthzPolicyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := apq.querySpec()
	_spec.Node.Columns = apq.fields
	if len(apq.fields) > 0 {
		_spec.Unique = apq.unique != nil && *apq.unique
	}
	return sqlgraph.CountNodes(ctx, apq.driver, _spec)
}

func (apq *AuthzPolicyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := apq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (apq *AuthzPolicyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authzpolicy.Table,
			Columns: authzpolicy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authzpolicy.FieldID,
			},
		},
		From:   apq.sql,
		Unique: true,
	}
	if unique := apq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := apq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authzpolicy.FieldID)
		for i := range fields {
			if fields[i] != authzpolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (apq *AuthzPolicyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(authzpolicy.Table)
	columns := apq.fields
	if len(columns) == 0 {
		columns = authzpolicy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if apq.unique != nil && *apq.unique {
		selector.Distinct()
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector)
	}
	if offset := apq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthzPolicyGroupBy is the group-by builder for AuthzPolicy entities.
type AuthzPolicyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *AuthzPolicyGroupBy) Aggregate(fns ...AggregateFunc) *AuthzPolicyGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the group-by query and scans the result into the given value.
func (apgb *AuthzPolicyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := apgb.path(ctx)
	if err != nil {
		return err
	}
	apgb.sql = query
	return apgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := apgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) StringsX(ctx context.Context) []string {
	v, err := apgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = apgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) StringX(ctx context.Context) string {
	v, err := apgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) IntsX(ctx context.Context) []int {
	v, err := apgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = apgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) IntX(ctx context.Context) int {
	v, err := apgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := apgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = apgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := apgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := apgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AuthzPolicyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = apgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (apgb *AuthzPolicyGroupBy) BoolX(ctx context.Context) bool {
	v, err := apgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (apgb *AuthzPolicyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range apgb.fields {
		if !authzpolicy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := apgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (apgb *AuthzPolicyGroupBy) sqlQuery() *sql.Selector {
	selector := apgb.sql.Select()
	aggregation := make([]string, 0, len(apgb.fns))
	for _, fn := range apgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(apgb.fields)+len(apgb.fns))
		for _, f := range apgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(apgb.fields...)...)
}

// AuthzPolicySelect is the builder for selecting fields of AuthzPolicy entities.
type AuthzPolicySelect struct {
	*AuthzPolicyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aps *AuthzPolicySelect) Scan(ctx context.Context, v interface{}) error {
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	aps.sql = aps.AuthzPolicyQuery.sqlQuery(ctx)
	return aps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aps *AuthzPolicySelect) ScanX(ctx context.Context, v interface{}) {
	if err := aps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) Strings(ctx context.Context) ([]string, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aps *AuthzPolicySelect) StringsX(ctx context.Context) []string {
	v, err := aps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aps *AuthzPolicySelect) StringX(ctx context.Context) string {
	v, err := aps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) Ints(ctx context.Context) ([]int, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aps *AuthzPolicySelect) IntsX(ctx context.Context) []int {
	v, err := aps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aps *AuthzPolicySelect) IntX(ctx context.Context) int {
	v, err := aps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aps *AuthzPolicySelect) Float64sX(ctx context.Context) []float64 {
	v, err := aps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aps *AuthzPolicySelect) Float64X(ctx context.Context) float64 {
	v, err := aps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AuthzPolicySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aps *AuthzPolicySelect) BoolsX(ctx context.Context) []bool {
	v, err := aps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aps *AuthzPolicySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authzpolicy.Label}
	default:
		err = fmt.Errorf("ent: AuthzPolicySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aps *AuthzPolicySelect) BoolX(ctx context.Context) bool {
	v, err := aps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aps *AuthzPolicySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aps.sql.Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
