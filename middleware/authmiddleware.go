package middleware

import (
	"basic_api/util"
	"encoding/json"
	"net/http"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")

		if err != nil {
			json.NewEncoder(w).Encode("something went wrong")
			return
		}

		if _, err := util.ParseJwt(cookie.Value); err != nil {
			json.NewEncoder(w).Encode("not auhtorized")
			return
		}
		handler.ServeHTTP(w, r)
	}

}
