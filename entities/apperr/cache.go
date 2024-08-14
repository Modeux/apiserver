package apperr

import "net/http"

var CacheErr = func(err error) error {
	return NewAppErr("common-a5iiR1xE", "Something went wrong", err, http.StatusInternalServerError)
}
