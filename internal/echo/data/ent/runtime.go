// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/yangchnet/skeleton/internal/echo/data/ent/echo"
	"github.com/yangchnet/skeleton/internal/echo/data/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	echoFields := schema.Echo{}.Fields()
	_ = echoFields
	// echoDescMessage is the schema descriptor for message field.
	echoDescMessage := echoFields[0].Descriptor()
	// echo.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	echo.MessageValidator = echoDescMessage.Validators[0].(func(string) error)
	// echoDescDeleted is the schema descriptor for deleted field.
	echoDescDeleted := echoFields[2].Descriptor()
	// echo.DefaultDeleted holds the default value on creation for the deleted field.
	echo.DefaultDeleted = echoDescDeleted.Default.(bool)
	// echoDescCreateTime is the schema descriptor for create_time field.
	echoDescCreateTime := echoFields[3].Descriptor()
	// echo.DefaultCreateTime holds the default value on creation for the create_time field.
	echo.DefaultCreateTime = echoDescCreateTime.Default.(func() time.Time)
}
