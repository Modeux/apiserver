package apperr

import (
	"github.com/go-chi/render"
	"pornrangers/pkg/responder"
)

type AppErr struct {
	HttpStatusCode int
	Err            error
	ErrText        string
	ErrCode        string
}

func NewAppErr(errCode, ErrText string, err error, httpStatusCode int) error {
	return &AppErr{
		HttpStatusCode: httpStatusCode,
		ErrCode:        errCode,
		ErrText:        ErrText,
		Err:            err, // error with stacktrace
	}
}

func (a *AppErr) Error() string {
	return a.Err.Error()
}

func (a *AppErr) ToErrResp() render.Renderer {
	return responder.NewErrResp(a.ErrCode, a.ErrText, a.HttpStatusCode)
}
