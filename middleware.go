package sharedtimer

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/api/idtoken"
)

// ValidateJWT validates the JWT in the "Authorization" header. If the JWT
// is missing or invalid a 401 is returned and the request is terminated.
// If the JWT is valid the trimmed token is stored in the "Token-raw"
// header and the subject is stored in the "Subject" header.
func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		tokenElements := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(tokenElements) < 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		payload, err := idtoken.Validate(context.TODO(), tokenElements[1], "")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r.Header.Set("Token-raw", tokenElements[1])
		r.Header.Set("Subject", payload.Subject)
		next.ServeHTTP(w, r)
	})
}
