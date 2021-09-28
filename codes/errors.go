package codes

import "errors"

var (
	// 500系エラー
	ErrInternalServerError = errors.New("internal server error")

	// 400系エラー
	ErrUserNotFound                  = errors.New("user not found")
	ErrUserAlreadyExisted            = errors.New("user has already existed")
	ErrUserMailAddressAlreadyExisted = errors.New("email has already existed")
	ErrPasswordTooLong               = errors.New("password too long")
)
