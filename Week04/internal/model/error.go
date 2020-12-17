package model

import "errors"

var (
	ErrCreateUser = errors.New("create user failed")
	ErrDelUser    = errors.New("delete user failed")
)

func MapErr2Code(err error) (code int32, msg string) {
	switch err {
	case ErrCreateUser:
		return 1001, err.Error()
	case ErrDelUser:
		return 1002, err.Error()
	default:
		// 表示位置错误
		return 1000, ""
	}
}
