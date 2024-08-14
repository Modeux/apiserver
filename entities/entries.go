package entities

import (
	"net/http"
	"time"
)

type EntryHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type LoginSvcInterface interface {
	Login(email, password string) (LoginData, error)
}

type LoginRepoInterface interface {
	GetUserByEmail(email string) (UserLogin, error)
}

type RegisterSvcInterface interface {
	CheckEmail(email string) error
	InsertUser(data SignUpRequest) error
}

type RegisterRepoInterface interface {
	CheckEmail(email string) (bool, error)
	InsertUser(data SignUpData) error
}

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=255"`
	Email    string `json:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}

type SignUpData struct {
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func NewSignUpData(data SignUpRequest, encryptedPass string) SignUpData {
	sd := SignUpData{
		Name:      data.Name,
		Email:     data.Email,
		Password:  encryptedPass,
		CreatedAt: time.Now().Format(time.DateTime),
		UpdatedAt: time.Now().Format(time.DateTime),
	}
	return sd
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}

type LoginData struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

type UserLogin struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
