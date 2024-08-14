package apperr

import (
	"fmt"
	"net/http"
)

var MaxItemErr = func(err error, maxItem int64) error {
	text := fmt.Sprintf("exceed max items %d", maxItem)
	return NewAppErr("layout-m1jcoq1Lrg", text, err, http.StatusBadRequest)
}

var InvalidCategoryLangIdErr = func(err error, ids []int64) error {
	text := fmt.Sprintf("invalid category lang id %v", ids)
	return NewAppErr("layout-1Lrg", text, err, http.StatusBadRequest)
}
