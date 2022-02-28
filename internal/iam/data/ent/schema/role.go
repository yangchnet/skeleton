package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("rolename").NotEmpty().Unique(),
		field.String("status").NotEmpty().
			Match(regexp.MustCompile(`active|disable|delete`)).
			Comment("active or disable or delete"),
		field.Time("create_time").Immutable().Default(func() time.Time {
			return time.Now()
		}),
		field.Time("update_time").Optional().Nillable(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bindings", BindUserRole.Type),
	}
}

// Indexes of the Role.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("rolename").Unique(),
	}
}
