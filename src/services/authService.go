package services

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-keycloak/src/errors"
	"net/http"
	"os"
	"strings"

	"github.com/Nerzal/gocloak/v11"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var (
	clientId     string
	clientSecret string
	realm        string
	hostname     string
	port         string
)

var client gocloak.GoCloak

func InitializeOauthServer() {
	// assign env to var
	clientId = os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm = os.Getenv("KEYCLOAK_REALM")
	hostname = os.Getenv("KEYCLOAK_HOST")
	port = os.Getenv("KEYCLOAK_PORT")

	client = gocloak.NewClient(fmt.Sprintf("%s:%s", hostname, port), gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))
}

func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if len(authHeader) < 1 {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(errors.UnauthorizedError())
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]

		restyClient := client.RestyClient()
		restyClient.SetDebug(true)
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		ctx := context.Background()

		rptResult, err := client.RetrospectToken(ctx, accessToken, clientId, clientSecret, realm)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errors.BadRequestError(err.Error()))
			return
		}

		isTokenValid := *rptResult.Active

		if !isTokenValid {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(errors.UnauthorizedError())
			return
		}

		next.ServeHTTP(w, r)
	})
}
