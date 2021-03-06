package token

import (
	"time"
)

// Maker is a token maker.
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
