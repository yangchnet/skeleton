// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/yangchnet/skeleton/internal/iam/data/ent/authzpolicy"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/binduserrole"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/role"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/schema"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/tenant"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authzpolicyFields := schema.AuthzPolicy{}.Fields()
	_ = authzpolicyFields
	// authzpolicyDescPolicyName is the schema descriptor for policy_name field.
	authzpolicyDescPolicyName := authzpolicyFields[0].Descriptor()
	// authzpolicy.PolicyNameValidator is a validator for the "policy_name" field. It is called by the builders before save.
	authzpolicy.PolicyNameValidator = authzpolicyDescPolicyName.Validators[0].(func(string) error)
	// authzpolicyDescObj is the schema descriptor for obj field.
	authzpolicyDescObj := authzpolicyFields[1].Descriptor()
	// authzpolicy.ObjValidator is a validator for the "obj" field. It is called by the builders before save.
	authzpolicy.ObjValidator = authzpolicyDescObj.Validators[0].(func(string) error)
	// authzpolicyDescStatus is the schema descriptor for status field.
	authzpolicyDescStatus := authzpolicyFields[3].Descriptor()
	// authzpolicy.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	authzpolicy.StatusValidator = func() func(string) error {
		validators := authzpolicyDescStatus.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(status string) error {
			for _, fn := range fns {
				if err := fn(status); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// authzpolicyDescCreateTime is the schema descriptor for create_time field.
	authzpolicyDescCreateTime := authzpolicyFields[4].Descriptor()
	// authzpolicy.DefaultCreateTime holds the default value on creation for the create_time field.
	authzpolicy.DefaultCreateTime = authzpolicyDescCreateTime.Default.(func() time.Time)
	binduserroleFields := schema.BindUserRole{}.Fields()
	_ = binduserroleFields
	// binduserroleDescUsername is the schema descriptor for username field.
	binduserroleDescUsername := binduserroleFields[0].Descriptor()
	// binduserrole.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	binduserrole.UsernameValidator = binduserroleDescUsername.Validators[0].(func(string) error)
	// binduserroleDescRolename is the schema descriptor for rolename field.
	binduserroleDescRolename := binduserroleFields[1].Descriptor()
	// binduserrole.RolenameValidator is a validator for the "rolename" field. It is called by the builders before save.
	binduserrole.RolenameValidator = binduserroleDescRolename.Validators[0].(func(string) error)
	// binduserroleDescCreateTime is the schema descriptor for create_time field.
	binduserroleDescCreateTime := binduserroleFields[2].Descriptor()
	// binduserrole.DefaultCreateTime holds the default value on creation for the create_time field.
	binduserrole.DefaultCreateTime = binduserroleDescCreateTime.Default.(func() time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescRolename is the schema descriptor for rolename field.
	roleDescRolename := roleFields[0].Descriptor()
	// role.RolenameValidator is a validator for the "rolename" field. It is called by the builders before save.
	role.RolenameValidator = roleDescRolename.Validators[0].(func(string) error)
	// roleDescStatus is the schema descriptor for status field.
	roleDescStatus := roleFields[1].Descriptor()
	// role.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	role.StatusValidator = func() func(string) error {
		validators := roleDescStatus.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(status string) error {
			for _, fn := range fns {
				if err := fn(status); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// roleDescCreateTime is the schema descriptor for create_time field.
	roleDescCreateTime := roleFields[2].Descriptor()
	// role.DefaultCreateTime holds the default value on creation for the create_time field.
	role.DefaultCreateTime = roleDescCreateTime.Default.(func() time.Time)
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescTenantName is the schema descriptor for tenant_name field.
	tenantDescTenantName := tenantFields[0].Descriptor()
	// tenant.TenantNameValidator is a validator for the "tenant_name" field. It is called by the builders before save.
	tenant.TenantNameValidator = tenantDescTenantName.Validators[0].(func(string) error)
	// tenantDescStatus is the schema descriptor for status field.
	tenantDescStatus := tenantFields[1].Descriptor()
	// tenant.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	tenant.StatusValidator = func() func(string) error {
		validators := tenantDescStatus.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(status string) error {
			for _, fn := range fns {
				if err := fn(status); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tenantDescCreateTime is the schema descriptor for create_time field.
	tenantDescCreateTime := tenantFields[2].Descriptor()
	// tenant.DefaultCreateTime holds the default value on creation for the create_time field.
	tenant.DefaultCreateTime = tenantDescCreateTime.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPasswd is the schema descriptor for passwd field.
	userDescPasswd := userFields[2].Descriptor()
	// user.PasswdValidator is a validator for the "passwd" field. It is called by the builders before save.
	user.PasswdValidator = userDescPasswd.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[3].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[4].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[5].Descriptor()
	// user.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	user.StatusValidator = func() func(string) error {
		validators := userDescStatus.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(status string) error {
			for _, fn := range fns {
				if err := fn(status); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userFields[6].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
}