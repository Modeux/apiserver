package entries

import (
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"net/http"
	"pornrangers/entities"
	"pornrangers/entities/apperr"
	"pornrangers/pkg/databases"
	"pornrangers/pkg/loggers"
	"pornrangers/pkg/responder"
	"pornrangers/pkg/validators"
)

type EntryHandler struct {
	RegisterSvc entities.RegisterSvcInterface
	LoginSvc    entities.LoginSvcInterface
}

func NewEntryHandle(db databases.DBInterface) *EntryHandler {
	registerRepo := NewRegisterRepo(db)
	registerSvc := NewRegisterSvc(registerRepo)
	loginRepo := NewLoginRepo(db)
	loginSvc := NewLoginSvc(loginRepo)
	return &EntryHandler{registerSvc, loginSvc}
}

func (e *EntryHandler) Login(w http.ResponseWriter, r *http.Request) {
	var data entities.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		loggers.Logger.Errorf("%+v", errors.WithStack(err))
		responder.AppResponse(w, r, apperr.JsonErr(err).(*apperr.AppErr).ToErrResp())
		return
	}
	// validate the login data
	errFields, err := validators.Validate(data)
	if err != nil {
		loggers.Logger.Errorf("%+v", err.(*apperr.AppErr).Err)
		responder.AppResponse(w, r, apperr.ValidationErr(err).(*apperr.AppErr).ToErrResp())
		return
	}
	if errFields != nil {
		responder.AppResponse(w, r, responder.NewValidateErrResp("Validate error", errFields))
		return
	}
	// login
	loginData, err := e.LoginSvc.Login(data.Email, data.Password)
	if err != nil {
		loggers.Logger.Errorf("%+v", err.(*apperr.AppErr).Err)
		responder.AppResponse(w, r, err.(*apperr.AppErr).ToErrResp())
		return
	}
	// response JWT to client
	responder.AppResponse(w, r, responder.NewDataResp(http.StatusOK, loginData))
}

func (e *EntryHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var data entities.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		loggers.Logger.Errorf("%+v", errors.WithStack(err))
		responder.AppResponse(w, r, apperr.JsonErr(err).(*apperr.AppErr).ToErrResp())
		return
	}
	// validate post data
	errFields, err := validators.Validate(data)
	if err != nil {
		loggers.Logger.Errorf("%+v", err)
		responder.AppResponse(w, r, apperr.ValidationErr(err).(*apperr.AppErr).ToErrResp())
		return
	}
	if errFields != nil {
		responder.AppResponse(w, r, responder.NewValidateErrResp("Validate error", errFields))
		return
	}
	// check the email duplicate
	if err := e.RegisterSvc.CheckEmail(data.Email); err != nil {
		loggers.Logger.Errorf("%+v", err.(*apperr.AppErr).Err)
		responder.AppResponse(w, r, err.(*apperr.AppErr).ToErrResp())
		return
	}
	// insert user data to database
	if err := e.RegisterSvc.InsertUser(data); err != nil {
		loggers.Logger.Errorf("%+v", err)
		responder.AppResponse(w, r, err.(*apperr.AppErr).ToErrResp())
		return
	}
	responder.AppResponse(w, r, responder.NewSuccessResp("Sign up success!", http.StatusCreated))
}
