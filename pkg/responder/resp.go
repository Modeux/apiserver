package responder

import (
	"github.com/goccy/go-json"
	"net/http"
	"video-heatmap/entities/apperr"
)

const (
	ErrStatus = "error"
	OkStatus  = "ok"
)

type ErrResp struct {
	Err         error             `json:"-"`               // low-level runtime error
	StatusText  string            `json:"status"`          // user-level status message
	AppCode     string            `json:"code,omitempty"`  // application-specific error code
	ErrorText   string            `json:"error,omitempty"` // application-level error message, for debugging
	ErrorFields map[string]string `json:"errorFields,omitempty"`
}

func NewErrResp(w http.ResponseWriter, code, errText string, httpStatus int) {
	resp := ErrResp{
		StatusText:  ErrStatus,
		AppCode:     code,
		ErrorText:   errText,
		ErrorFields: map[string]string{},
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(resp)
}

func NewErrRespFromAppErr(w http.ResponseWriter, err error) {
	resp := ErrResp{
		StatusText: ErrStatus,
		AppCode:    err.(*apperr.AppErr).ErrCode,
		ErrorText:  err.(*apperr.AppErr).ErrText,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err.(*apperr.AppErr).HttpStatusCode)
	_ = json.NewEncoder(w).Encode(resp)
}

type ValidateErrResp struct {
	Err            error             `json:"-"`      // low-level runtime error
	HTTPStatusCode int               `json:"-"`      // http response status code
	StatusText     string            `json:"status"` // user-level status message
	ErrorText      string            `json:"error"`
	ErrorFields    map[string]string `json:"fields,omitempty"`
}

func NewValidateErrResp(w http.ResponseWriter, errText string, fields map[string]string) {
	resp := ValidateErrResp{
		StatusText:     ErrStatus,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		ErrorFields:    fields,
		ErrorText:      errText,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnprocessableEntity)
	_ = json.NewEncoder(w).Encode(resp)
}

type SuccessResp struct {
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	Message        string `json:"message"`
}

func NewSuccessResp(w http.ResponseWriter, msg string, httpStatus int) {
	resp := SuccessResp{
		HTTPStatusCode: httpStatus,
		Message:        msg,
		StatusText:     OkStatus,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(resp)
}

type DataResp[T any] struct {
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	Message        string `json:"message,omitempty"`
	Data           T      `json:"data"`
	ErrorText      string `json:"error,omitempty"`
	AppCode        string `json:"code,omitempty"`
	FirstId        int64  `json:"firstId,omitempty"`
	LastId         int64  `json:"lastId,omitempty"`
	Page           int64  `json:"page,omitempty"`
	Total          int64  `json:"total,omitempty"`
}

func NewDataResp[T any](w http.ResponseWriter, httpStatus int, data T) {
	resp := DataResp[T]{
		HTTPStatusCode: httpStatus,
		StatusText:     OkStatus,
		Data:           data,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(resp)
}

type NullResp struct {
	HTTPStatusCode int      `json:"-"`
	StatusText     string   `json:"status"`
	Data           []string `json:"data"`
}

func NewNullResp(httpStatus int) *NullResp {
	return &NullResp{
		HTTPStatusCode: httpStatus,
		StatusText:     OkStatus,
		Data:           nil,
	}
}
