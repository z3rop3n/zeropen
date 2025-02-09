package spector

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zeropen/app/sazs"
	"github.com/zeropen/app/spector/config"
	"github.com/zeropen/app/spector/token"
	"github.com/zeropen/app/spector/user"
)

// LoggingMiddleware wraps the request handler with logging
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println(r.Method)

		// Log the request details
		log.Printf(
			"Started %s %s from %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		// Log the completion time
		log.Printf(
			"Completed %s %s in %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

func Run(config sazs.Config, appConfig config.AppConfig) {
	r := mux.NewRouter()
	r.Use(LoggingMiddleware)
	r.HandleFunc("/", HomeHandler)
	tokenObj := token.NewTokenObj(appConfig.JWT_AUTH_SECRET)
	userSubrouter := r.PathPrefix("/user").Subrouter()
	tokenSubrouter := r.PathPrefix("/auth").Subrouter()
	user.NewUserAPI(&config, appConfig, *tokenObj).RegisterRouters(userSubrouter, config)
	token.NewTokenAPI(&config, appConfig).RegisterRouters(tokenSubrouter, config)

	log.Println("Starting server on :8080")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(r)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Gorilla!"))
}
