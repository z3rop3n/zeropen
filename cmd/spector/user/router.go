package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeropen/app/sazs"
	"github.com/zeropen/app/spector/config"
)

type UserAPI struct {
	config    *sazs.Config
	appConfig config.AppConfig
}

func NewUserAPI(config *sazs.Config, appConfig config.AppConfig) *UserAPI {
	return &UserAPI{config: config, appConfig: appConfig}
}

func (uApi *UserAPI) RegisterRouters(r *mux.Router, config sazs.Config) *mux.Router {
	r.HandleFunc("/signup", uApi.SignupHandler).Methods("POST")
	r.HandleFunc("/verifyotp", uApi.VerifyOTPHandler).Methods("POST")

	return r
}

func (uApi *UserAPI) SignupHandler(w http.ResponseWriter, r *http.Request) {
	type SignupRequest struct {
		Email string `json:"email"`
	}
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type SignupResponse struct {
	}
	var resp SignupResponse
	code, err := uApi.Signup(req.Email)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uApi *UserAPI) VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	type VerifyOTPRequest struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	var req VerifyOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type VerifyOTPResponse struct {
	}
	var resp VerifyOTPResponse
	code, err := uApi.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
