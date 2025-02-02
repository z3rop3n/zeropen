package token

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeropen/app/sazs"
	"github.com/zeropen/app/spector/config"
)

type TokenAPI struct {
	config    *sazs.Config
	appConfig config.AppConfig
}

func NewTokenAPI(config *sazs.Config, appConfig config.AppConfig) *TokenAPI {
	return &TokenAPI{config: config, appConfig: appConfig}
}

func (tApi *TokenAPI) RegisterRouters(r *mux.Router, config sazs.Config) *mux.Router {
	r.HandleFunc("/refresh", tApi.RefreshHandler).Methods("POST")
	return r
}

func (tApi *TokenAPI) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	type RefreshRequest struct {
		RefreshToken string `json:"refreshToken"`
	}
	var req RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	code, resp, err := tApi.RefreshAccessToken(req.RefreshToken)
	type RefreshResponse struct {
		AccessToken *string `json:"accessToken"`
	}
	var res RefreshResponse
	res.AccessToken = resp
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
