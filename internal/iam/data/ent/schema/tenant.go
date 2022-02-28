package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("tenant_name").NotEmpty().Unique(),
		field.String("status").NotEmpty().
			Match(regexp.MustCompile(`active|disable|delete`)).
			Comment("active or disable or delete"),
		field.Time("create_time").Immutable().Default(func() time.Time {
			return time.Now()
		}),
		field.Time("update_time").Optional().Nillable(),
	}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}

// Indexes of the Tenant.
func (Tenant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_name").Unique(),
	}
}
