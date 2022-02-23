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
		field.Time("create_time").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("update_time"),
		field.Time("delete_time"),
	}
}

// Edges of the Echo.
func (Echo) Edges() []ent.Edge {
	return nil
}
