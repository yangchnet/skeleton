package passwd

import (
	"testing"

	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
)

func TestPassword(t *testing.T) {
	// passwd := "admin"
	// crypto_passwd, err := HashPassword(passwd)
	// require.NoError(t, err)
	// // fmt.Println(crypto_passwd)

	// err = CheckPassword(passwd, crypto_passwd)
	// require.NoError(t, err)

	p := ladon.DefaultPolicy{
		ID:          "1",
		Description: "root policy",
		Subjects:    []string{"root"},
		Effect:      ladon.AllowAccess,
		Resources:   []string{"system<.*>"},
		Actions:     []string{"delete"},
		Meta:        []byte{},
	}

	req := ladon.Request{
		Resource: "system",
		Action:   "delete",
		Subject:  "root",
	}

	warden := &ladon.Ladon{Manager: memory.NewMemoryManager()}
	warden.Manager.Create(&p)
	if err := warden.IsAllowed(&req); err == nil {
		t.Log("授权成功")
		return
	}

	t.Log("false")
}
