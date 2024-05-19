package errs

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrNIPAlreadyRegistered = errors.New("nip already registered")
var ErrInvalidNIP = errors.New("nip is invalid")
var ErrInvalidPassword = errors.New("invalid password")

var ErrUserITNotFound = errors.New("user IT not found")

var ErrUserNurseNotFound = errors.New("user nurse not found")
