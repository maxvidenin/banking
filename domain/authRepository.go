package domain

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/maxvidenin/banking/logger"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, routeVars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (ra RemoteAuthRepository) IsAuthorized(token string, routeName string, routeVars map[string]string) bool {
	u := buildVerifyUrl(token, routeName, routeVars)
	if res, err := http.Get(u); err != nil {
		logger.Error(err.Error())
		return false
	} else {
		m := map[string]bool{}
		// TODO fix error: "json: cannot unmarshal string into Go value of type bool"
		if err = json.NewDecoder(res.Body).Decode(&m); err != nil {
			logger.Error(err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

func buildVerifyUrl(token string, routeName string, routeVars map[string]string) string {
	u := url.URL{Host: "localhost:8181", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Set("token", token)
	q.Set("routeName", routeName)
	for k, v := range routeVars {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
