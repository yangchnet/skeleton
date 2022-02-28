package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.String("tenant_name").Optional().Nillable(),
		field.String("passwd").NotEmpty(),
		field.String("phone").Optional().Nillable().Match(regexp.MustCompile(`\d{11}`)),
		field.String("email").Optional().Nillable().Match(regexp.MustCompile(`.*@\w*\.com`)),
		field.String("status").NotEmpty().
			Match(regexp.MustCompile(`active|disable|delete`)).
			Comment("active or disable or delete"),
		field.Time("create_time").Immutable().Default(func() time.Time {
			return time.Now()
		}),
		field.Time("update_time").Optional().Nillable(),
	}
}

// Indexes of user.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "passwd").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("policys", AuthzPolicy.Type),
		edge.To("bindings", BindUserRole.Type),

		edge.From("belong", Tenant.Type).Ref("users").Unique(),
	}
}
