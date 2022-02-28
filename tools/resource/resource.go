package resource

import "strings"

// Build build resource ID from given string.
func Build(tenantName, resource string) string {
	return strings.Join([]string{"system", tenantName, resource}, "::")
}
