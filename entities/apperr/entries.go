package apperr

import (
	"errors"
	"net/http"
)

var EmailExistErr = errors.New("email already exist")

var LoginErr = func(err error) error {
	return NewAppErr("Login-ZchOhZgo", "Email or password error", err, http.StatusBadRequest)
}

var CreateTokenErr = func(err error) error {
	return NewAppErr("Login-94I3rw", "Could not create the login credentials", err, http.StatusBadRequest)
}

var PasswordEncryptErr = func(err error) error {
	return NewAppErr("Signup-N7n", "Password encryption error", err, http.StatusBadRequest)
}

var SignUpEmailDupErr = func(err error) error {
	return NewAppErr("Signup-1v3jBs4", err.Error(), err, http.StatusBadRequest)
}
