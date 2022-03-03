package biz

import (
	"context"
	"time"

	"github.com/yangchnet/skeleton/pkg/logger"
	"github.com/yangchnet/skeleton/tools/passwd"
)

type Token struct {
	Jwt string
}

func (uc *IamCase) CreateToken(ctx context.Context, username, password string) (*Token, error) {
	pw, err := uc.userRepo.Passwd(ctx, username)
	if err != nil {
		return nil, err
	}

	if err := passwd.CheckPassword(password, pw); err != nil {
		logger.Error("check password failed ", err)
		return nil, err
	}

	token, err := uc.tokenMaker.CreateToken(username, time.Second*7200)
	if err != nil {
		return nil, err
	}

	return &Token{
		Jwt: token,
	}, nil
}
