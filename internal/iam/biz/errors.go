package biz

import "errors"

var (
	ErrExistUsername = errors.New("已存在的用户名")
	ErrNotFound      = errors.New("记录不存在")
)
