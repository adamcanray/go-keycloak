package main

import (
	"go-keycloak/src/config"
	"go-keycloak/src/controllers"
	"go-keycloak/src/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file only on local environtment
	if os.Getenv("MODE") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	run()
}

func run() {
	config.DbConnect()

	services.InitializeOauthServer()

	router := mux.NewRouter().StrictSlash(true)

	// use middleware globally
	router.Use(commonMiddleware)

	registerRoutes(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func registerRoutes(router *mux.Router) {
	registerControllerRoutes(controllers.EventController{}, router)
}

func registerControllerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}

// a middleware that sets the Content-Type header for all request to application/json
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
