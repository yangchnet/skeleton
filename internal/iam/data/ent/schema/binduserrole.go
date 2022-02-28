package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BindUserRole holds the schema definition for the BindUserRole entity.
type BindUserRole struct {
	ent.Schema
}

// Fields of the BindUserRole.
func (BindUserRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty(),
		field.String("rolename").NotEmpty(),
		field.Time("create_time").Immutable().Default(func() time.Time {
			return time.Now()
		}),
		field.Time("update_time").Optional().Nillable(),
	}
}

// Edges of the BindUserRole.
func (BindUserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("bindings").Unique(),
		edge.From("role", Role.Type).Ref("bindings").Unique(),
	}
}

// Indexes of the BindUserRole.
func (BindUserRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "rolename"),
	}
}
