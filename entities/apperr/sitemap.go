package apperr

import "net/http"

var SitemapErr = func(err error) error {
	return NewAppErr("sitemap-As1pqAG", "Get sitemap error", err, http.StatusBadRequest)
}
