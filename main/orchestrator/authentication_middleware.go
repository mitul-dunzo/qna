package orchestrator

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"qna/main/services"
	"strings"
)

type AuthenticationMiddleware struct {
	JwtService services.JwtService
}

var openApis = []string{
	"/auth",
}

func NewAuthenticationMiddleware(jwtService services.JwtService) AuthenticationMiddleware {
	return AuthenticationMiddleware{JwtService: jwtService}
}

func (orch *AuthenticationMiddleware) Check(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if orch.isOpen(r) {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		if len(token) < 5 {
			http.Error(w, "bad auth token", http.StatusForbidden)
			return
		}

		if token[0:4] != "JWT " {
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}
		token = token[4:]
		userId, err := orch.JwtService.ValidateUser(token)
		if err != nil {
			logrus.Error("Unauthorized request")
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		logrus.Debug("userID: ", userId)

		ctx := context.WithValue(r.Context(), "user_id", userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (orch *AuthenticationMiddleware) isOpen(r *http.Request) bool {
	urlPath := r.URL.Path
	for i := 0; i < len(openApis); i++ {
		if strings.HasPrefix(urlPath, openApis[i]) {
			return true
		}
	}
	return false
}
