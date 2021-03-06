// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthzPoliciesColumns holds the columns for the "authz_policies" table.
	AuthzPoliciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "policy_name", Type: field.TypeString},
		{Name: "obj", Type: field.TypeString},
		{Name: "policy", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "status", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime, Nullable: true},
		{Name: "user_policys", Type: field.TypeInt, Nullable: true},
	}
	// AuthzPoliciesTable holds the schema information for the "authz_policies" table.
	AuthzPoliciesTable = &schema.Table{
		Name:       "authz_policies",
		Columns:    AuthzPoliciesColumns,
		PrimaryKey: []*schema.Column{AuthzPoliciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "authz_policies_users_policys",
				Columns:    []*schema.Column{AuthzPoliciesColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "authzpolicy_obj",
				Unique:  false,
				Columns: []*schema.Column{AuthzPoliciesColumns[2]},
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rolename", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime, Nullable: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_rolename",
				Unique:  true,
				Columns: []*schema.Column{RolesColumns[1]},
			},
		},
	}
	// TenantsColumns holds the columns for the "tenants" table.
	TenantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tenant_name", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime, Nullable: true},
	}
	// TenantsTable holds the schema information for the "tenants" table.
	TenantsTable = &schema.Table{
		Name:       "tenants",
		Columns:    TenantsColumns,
		PrimaryKey: []*schema.Column{TenantsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tenant_tenant_name",
				Unique:  true,
				Columns: []*schema.Column{TenantsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "tenant_name", Type: field.TypeString, Nullable: true},
		{Name: "passwd", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime, Nullable: true},
		{Name: "tenant_users", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_tenants_users",
				Columns:    []*schema.Column{UsersColumns[9]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_passwd",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[3]},
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthzPoliciesTable,
		RolesTable,
		TenantsTable,
		UsersTable,
		UserRolesTable,
	}
)

func init() {
	AuthzPoliciesTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = TenantsTable
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
}
