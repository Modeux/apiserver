package apperr

import "net/http"

var BuildErr = func(err error) error {
	return NewAppErr("build-Gz5DCmb4", "Build frontend error", err, http.StatusInternalServerError)
}
