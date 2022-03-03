package biz

import (
	"errors"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewIamCase)

var (
	ErrExistUsername = errors.New("已存在的用户名")
	ErrNotFound      = errors.New("记录不存在")
	ErrNotFoundUser  = errors.New("用户名不存在")
)
