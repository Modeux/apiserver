package auth

import (
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/pkg/errors"
	"os"
)

func GetJwtPublicKey() (jwk.Key, error) {
	keySet, err := os.ReadFile(os.Getenv("APP_RSA_KEY"))
	if err != nil {
		return nil, errors.Wrap(err, "Read jwt key file")
	}
	privateKey, err := jwk.ParseKey(keySet)
	if err != nil {
		return nil, errors.Wrap(err, "Parse jwt private key")
	}
	pubKey, err := jwk.PublicKeyOf(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "Parse jwt public key")
	}
	return pubKey, nil
}
