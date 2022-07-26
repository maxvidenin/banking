package app

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/maxvidenin/banking-lib/errs"
	"github.com/maxvidenin/banking/domain"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (am AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			curtentRoute := mux.CurrentRoute(r)
			curtentRouteVars := mux.Vars(r)

			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthotized := am.repo.IsAuthorized(token, curtentRoute.GetName(), curtentRouteVars)

				if isAuthotized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(w, http.StatusUnauthorized, "Token is not provided")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
