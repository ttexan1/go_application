package web

import (
	"context"
	"net/http"
	"strings"

	"github.com/ttexan1/golang-simple/domain"
)

type ctxkey int

const (
	ctxClaims ctxkey = iota
)

func withAuth(h http.Handler, required bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := jwtFromAuth(r.Header.Get("Authorization"))
		if err == nil {
			ctx := context.WithValue(r.Context(), ctxClaims, claims)
			h.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		if required {
			sendErrorJSON(w, err)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func jwtFromAuth(auth string) (claims *domain.JWTClaims, err *domain.Error) {
	if auth == "" {
		err = domain.NewError(http.StatusUnauthorized, "missing Authorization header")
		return
	}
	if len(auth) < 7 {
		err = domain.NewError(http.StatusUnauthorized, "invalid Authorization header")
		return
	}
	if strings.ToUpper(auth[0:7]) != "BEARER " {
		err = domain.NewError(http.StatusUnauthorized, "Authorization header must start with bearer ")
		return
	}
	decrypted, err := domain.DecryptJWTToken(auth[7:])
	if err != nil {
		return
	}
	return decrypted, nil
}
