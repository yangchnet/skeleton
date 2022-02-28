package resource

import "testing"

func Test_Build(t *testing.T) {
	t.Log(Build("tenant1", "article"))
}
