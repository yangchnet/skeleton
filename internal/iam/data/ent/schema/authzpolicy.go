package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AuthzPolicy holds the schema definition for the AuthzPolicy entity.
type AuthzPolicy struct {
	ent.Schema
}

// Fields of the AuthzPolicy.
func (AuthzPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.String("policy_name").NotEmpty(),
		field.String("obj").NotEmpty(),
		field.Text("policy").Optional().Nillable(),
		field.String("status").NotEmpty().
			Match(regexp.MustCompile(`active|disable|delete`)).
			Comment("active or disable or delete"),
		field.Time("create_time").Immutable().Default(func() time.Time {
			return time.Now()
		}),
		field.Time("update_time").Optional().Nillable(),
	}
}

// Edges of the AuthzPolicy.
func (AuthzPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("create_by", User.Type).Ref("policys").Unique(),
	}
}

// Indexes of the AuthzPolicy.
func (AuthzPolicy) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("obj"),
	}
}
