// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EchosColumns holds the columns for the "echos" table.
	EchosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "message", Type: field.TypeString},
		{Name: "echo_message", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime},
	}
	// EchosTable holds the schema information for the "echos" table.
	EchosTable = &schema.Table{
		Name:       "echos",
		Columns:    EchosColumns,
		PrimaryKey: []*schema.Column{EchosColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EchosTable,
	}
)

func init() {
}
