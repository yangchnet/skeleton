// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/yangchnet/skeleton/internal/echo/data/ent/echo"
	"github.com/yangchnet/skeleton/internal/echo/data/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEcho = "Echo"
)

// EchoMutation represents an operation that mutates the Echo nodes in the graph.
type EchoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	message       *string
	echo_message  *string
	create_time   *time.Time
	update_time   *time.Time
	delete_time   *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Echo, error)
	predicates    []predicate.Echo
}

var _ ent.Mutation = (*EchoMutation)(nil)

// echoOption allows management of the mutation configuration using functional options.
type echoOption func(*EchoMutation)

// newEchoMutation creates new mutation for the Echo entity.
func newEchoMutation(c config, op Op, opts ...echoOption) *EchoMutation {
	m := &EchoMutation{
		config:        c,
		op:            op,
		typ:           TypeEcho,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEchoID sets the ID field of the mutation.
func withEchoID(id int) echoOption {
	return func(m *EchoMutation) {
		var (
			err   error
			once  sync.Once
			value *Echo
		)
		m.oldValue = func(ctx context.Context) (*Echo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Echo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEcho sets the old Echo of the mutation.
func withEcho(node *Echo) echoOption {
	return func(m *EchoMutation) {
		m.oldValue = func(context.Context) (*Echo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EchoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EchoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EchoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EchoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Echo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetMessage sets the "message" field.
func (m *EchoMutation) SetMessage(s string) {
	m.message = &s
}

// Message returns the value of the "message" field in the mutation.
func (m *EchoMutation) Message() (r string, exists bool) {
	v := m.message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "message" field's value of the Echo entity.
// If the Echo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EchoMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "message" field.
func (m *EchoMutation) ResetMessage() {
	m.message = nil
}

// SetEchoMessage sets the "echo_message" field.
func (m *EchoMutation) SetEchoMessage(s string) {
	m.echo_message = &s
}

// EchoMessage returns the value of the "echo_message" field in the mutation.
func (m *EchoMutation) EchoMessage() (r string, exists bool) {
	v := m.echo_message
	if v == nil {
		return
	}
	return *v, true
}

// OldEchoMessage returns the old "echo_message" field's value of the Echo entity.
// If the Echo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EchoMutation) OldEchoMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEchoMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEchoMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEchoMessage: %w", err)
	}
	return oldValue.EchoMessage, nil
}

// ResetEchoMessage resets all changes to the "echo_message" field.
func (m *EchoMutation) ResetEchoMessage() {
	m.echo_message = nil
}

// SetCreateTime sets the "create_time" field.
func (m *EchoMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *EchoMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Echo entity.
// If the Echo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EchoMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *EchoMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *EchoMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *EchoMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Echo entity.
// If the Echo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EchoMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *EchoMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetDeleteTime sets the "delete_time" field.
func (m *EchoMutation) SetDeleteTime(t time.Time) {
	m.delete_time = &t
}

// DeleteTime returns the value of the "delete_time" field in the mutation.
func (m *EchoMutation) DeleteTime() (r time.Time, exists bool) {
	v := m.delete_time
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteTime returns the old "delete_time" field's value of the Echo entity.
// If the Echo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EchoMutation) OldDeleteTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeleteTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeleteTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteTime: %w", err)
	}
	return oldValue.DeleteTime, nil
}

// ResetDeleteTime resets all changes to the "delete_time" field.
func (m *EchoMutation) ResetDeleteTime() {
	m.delete_time = nil
}

// Where appends a list predicates to the EchoMutation builder.
func (m *EchoMutation) Where(ps ...predicate.Echo) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *EchoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Echo).
func (m *EchoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EchoMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.message != nil {
		fields = append(fields, echo.FieldMessage)
	}
	if m.echo_message != nil {
		fields = append(fields, echo.FieldEchoMessage)
	}
	if m.create_time != nil {
		fields = append(fields, echo.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, echo.FieldUpdateTime)
	}
	if m.delete_time != nil {
		fields = append(fields, echo.FieldDeleteTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EchoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case echo.FieldMessage:
		return m.Message()
	case echo.FieldEchoMessage:
		return m.EchoMessage()
	case echo.FieldCreateTime:
		return m.CreateTime()
	case echo.FieldUpdateTime:
		return m.UpdateTime()
	case echo.FieldDeleteTime:
		return m.DeleteTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EchoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case echo.FieldMessage:
		return m.OldMessage(ctx)
	case echo.FieldEchoMessage:
		return m.OldEchoMessage(ctx)
	case echo.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case echo.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case echo.FieldDeleteTime:
		return m.OldDeleteTime(ctx)
	}
	return nil, fmt.Errorf("unknown Echo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EchoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case echo.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	case echo.FieldEchoMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEchoMessage(v)
		return nil
	case echo.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case echo.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case echo.FieldDeleteTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteTime(v)
		return nil
	}
	return fmt.Errorf("unknown Echo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EchoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EchoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EchoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Echo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EchoMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EchoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EchoMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Echo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EchoMutation) ResetField(name string) error {
	switch name {
	case echo.FieldMessage:
		m.ResetMessage()
		return nil
	case echo.FieldEchoMessage:
		m.ResetEchoMessage()
		return nil
	case echo.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case echo.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case echo.FieldDeleteTime:
		m.ResetDeleteTime()
		return nil
	}
	return fmt.Errorf("unknown Echo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EchoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EchoMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EchoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EchoMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EchoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EchoMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EchoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Echo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EchoMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Echo edge %s", name)
}
