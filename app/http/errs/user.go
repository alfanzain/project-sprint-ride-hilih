package errs

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrNIPAlreadyRegistered = errors.New("nip already registered")
var ErrUserITNotFound = errors.New("user IT not found")
var ErrInvalidPassword = errors.New("invalid password")
var ErrInvalidNIP = errors.New("nip is invalid")
var ErrNotNurse = errors.New("this user is not nurse")

