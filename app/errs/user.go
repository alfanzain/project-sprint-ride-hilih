package errs

import "errors"

var ErrNIPAlreadyRegistered = errors.New("nip already registered")
var ErrUserITNotFound = errors.New("user IT not found")
var ErrInvalidPassword = errors.New("invalid password")
var ErrInvalidNIP = errors.New("nip is invalid")
