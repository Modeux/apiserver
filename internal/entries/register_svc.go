package entries

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"pornrangers/entities"
	"pornrangers/entities/apperr"
)

type RegisterSvc struct {
	EntryRepo entities.RegisterRepoInterface
}

func NewRegisterSvc(entryRepo entities.RegisterRepoInterface) *RegisterSvc {
	return &RegisterSvc{entryRepo}
}

func (e *RegisterSvc) CheckEmail(email string) error {
	exist, err := e.EntryRepo.CheckEmail(email)
	if err != nil {
		return apperr.DBErr(err)
	}
	if exist {
		return apperr.SignUpEmailDupErr(errors.WithStack(apperr.EmailExistErr))
	}
	return nil
}

func (e *RegisterSvc) InsertUser(data entities.SignUpRequest) error {
	pass, err := e.encryptPassword(data.Password)
	if err != nil {
		return apperr.PasswordEncryptErr(err)
	}
	signupData := entities.NewSignUpData(data, pass)
	if err := e.EntryRepo.InsertUser(signupData); err != nil {
		return apperr.DBErr(err)
	}
	return nil
}

func (e *RegisterSvc) encryptPassword(password string) (string, error) {
	pass := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "encrypt password")
	}
	return string(hashedPassword), nil
}
