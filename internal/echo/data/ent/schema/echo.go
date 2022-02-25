package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Echo holds the schema definition for the Echo entity.
type Echo struct {
	ent.Schema
}

// Fields of the Echo.
func (Echo) Fields() []ent.Field {
	return []ent.Field{
		field.String("message").NotEmpty(),
		field.String("echo_message"),
		field.Bool("deleted").Default(false),
		field.Time("create_time").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("update_time").Nillable().Optional(),
		field.Time("delete_time").Nillable().Optional(),
	}
}

// Edges of the Echo.
func (Echo) Edges() []ent.Edge {
	return nil
}
