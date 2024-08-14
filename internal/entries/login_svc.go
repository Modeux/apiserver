package entries

import (
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"os"
	"pornrangers/entities"
	"pornrangers/entities/apperr"
	"strconv"
	"time"
)

type LoginSvc struct {
	LoginRepo entities.LoginRepoInterface
}

func NewLoginSvc(loginRepo entities.LoginRepoInterface) *LoginSvc {
	return &LoginSvc{LoginRepo: loginRepo}
}

func (l *LoginSvc) Login(email, password string) (entities.LoginData, error) {
	var data entities.LoginData
	user, err := l.LoginRepo.GetUserByEmail(email)
	if err != nil {
		return data, apperr.DBErr(err)
	}
	if user.Id == 0 {
		return data, apperr.LoginErr(errors.New("user not found"))
	}
	if err := l.ComparePassword(password, user.Password); err != nil {
		return data, apperr.LoginErr(err)
	}
	token, err := l.CreateAccessToken(user)
	if err != nil {
		return data, apperr.CreateTokenErr(err)
	}
	data.Email = email
	data.Name = user.Name
	data.AccessToken = token
	return data, nil
}

func (l *LoginSvc) ComparePassword(password, hashedPassword string) error {
	hp := []byte(hashedPassword)
	p := []byte(password)
	if err := bcrypt.CompareHashAndPassword(hp, p); err != nil {
		return errors.Wrap(err, "compare password")
	}
	return nil
}

func (l *LoginSvc) CreateAccessToken(user entities.UserLogin) (string, error) {
	buf, err := os.ReadFile(os.Getenv("APP_RSA_KEY"))
	if err != nil {
		return "", errors.Wrap(err, "create access token")
	}
	key, err := jwk.ParseKey(buf)
	if err != nil {
		return "", errors.Wrap(err, "create access token")
	}
	token, err := jwt.NewBuilder().
		Issuer(os.Getenv("APP_NAME")).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(168 * time.Hour)).
		Subject(strconv.FormatInt(user.Id, 10)).
		JwtID(uuid.NewString()).
		Build()
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, key))
	if err != nil {
		return "", errors.Wrap(err, "create access token")
	}
	return string(signed), nil
}
