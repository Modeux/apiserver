package apperr

import "net/http"

var DBErr = func(err error) error {
	return NewAppErr("common-ONT4dmD", "Something went wrong", err, http.StatusInternalServerError)
}

var ReqErr = func(err error) error {
	return NewAppErr("common-QEB0qk56", "Parse the request error", err, http.StatusBadRequest)
}

var ValidationErr = func(err error) error {
	return NewAppErr("common-xQottV3", "Invalid data", err, http.StatusInternalServerError)
}

var PageListConditionErr = func(err error) error {
	return NewAppErr("common-8Ob619", "Parse list condition error ", err, http.StatusBadRequest)
}

var JsonErr = func(err error) error {
	return NewAppErr("common-tl7i8", "Cloud not handle the json data", err, http.StatusBadRequest)
}

var StrconvErr = func(err error) error {
	return NewAppErr("common-0Y6CFHe0", "Cloud parse string to integer", err, http.StatusBadRequest)
}

var NotFoundErr = func(err error) error {
	return NewAppErr("common-DoZ", "Data not found", err, http.StatusNotFound)
}

var FileSizeErr = func(err error) error {
	return NewAppErr("common-C0B38", "File size is too large", err, http.StatusBadRequest)
}

var ReadFileErr = func(err error) error {
	return NewAppErr("common-Yc2c0NV", "Could not read file from request", err, http.StatusBadRequest)
}

var TimeParseErr = func(err error) error {
	return NewAppErr("common-UvG", "Could not parse the time format from request", err, http.StatusBadRequest)
}

var UploadImageErr = func(err error) error {
	return NewAppErr("common-6lZ", "Could not upload image", err, http.StatusBadRequest)
}

var CaseNotFoundErr = func(err error) error {
	return NewAppErr("common-cWd470rH1LG", "Case not found", err, http.StatusBadRequest)
}
