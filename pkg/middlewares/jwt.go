package middlewares

import (
	"apiserver/pkg/auth"
	"apiserver/pkg/responder"
	"context"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pubKey, err := auth.GetJwtPublicKey()
		if err != nil {
			responder.NewErrResp(w, "auth-001", "Auth error", http.StatusBadRequest)
			return
		}
		token, err := jwt.ParseRequest(r, jwt.WithKey(jwa.RS256, pubKey))
		if err != nil {
			responder.NewErrResp(w, "auth-002", "Auth fail", http.StatusBadRequest)
			return
		}
		if err := jwt.Validate(token); err != nil {
			responder.NewErrResp(w, "auth-003", "Auth fail", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "accessToken", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
