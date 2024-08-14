package apperr

import (
	"net/http"
)

var SlugDuplicateErrText = "Slug is already taken"
var SlugDuplicateErr = func(err error) error {
	return NewAppErr("site-As1pqAG", err.Error(), err, http.StatusBadRequest)
}
