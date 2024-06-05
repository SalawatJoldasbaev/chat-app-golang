package constants

import "errors"

var (
	ErrInvalidAuth   = errors.New("username or password is wrong")
	ErrIdRequired    = errors.New("id params is required")
	ErrEmailExist    = errors.New("email address already exist")
	ErrUsernameExist = errors.New("username already exist")
	ErrNameExist     = errors.New("name already exist")
)
