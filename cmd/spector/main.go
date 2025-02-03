package spector

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeropen/app/sazs"
	"github.com/zeropen/app/spector/config"
	"github.com/zeropen/app/spector/token"
	"github.com/zeropen/app/spector/user"
)
	
func Run(config sazs.Config, appConfig config.AppConfig) {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	tokenObj := token.NewTokenObj(appConfig.JWT_AUTH_SECRET)
	userSubrouter := r.PathPrefix("/user").Subrouter()
	user.NewUserAPI(&config, appConfig, *tokenObj).RegisterRouters(userSubrouter, config)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Gorilla!"))
}
