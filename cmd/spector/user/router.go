package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeropen/app/sazs"
	"github.com/zeropen/app/spector/config"
	"github.com/zeropen/app/spector/token"
	"github.com/zeropen/pkg/types"
)

type UserAPI struct {
	config          *sazs.Config
	appConfig       config.AppConfig
	tokenController *token.TokenController
}

func NewUserAPI(config *sazs.Config, appConfig config.AppConfig, tokenObj token.TokenController) *UserAPI {
	return &UserAPI{config: config, appConfig: appConfig, tokenController: &tokenObj}
}

func (uApi *UserAPI) RegisterRouters(r *mux.Router, config sazs.Config) *mux.Router {
	tokenController := *uApi.tokenController
	r.HandleFunc("/signup", uApi.SignupHandler).Methods("POST")
	r.HandleFunc("/verifyotp", uApi.VerifyOTPHandler).Methods("POST")
	r.HandleFunc("/profile", tokenController.AccessAuthMiddleware(http.HandlerFunc(uApi.GetUserProfileHandler))).Methods("GET")
	r.HandleFunc("/update", tokenController.AccessAuthMiddleware(http.HandlerFunc(uApi.UpdateHandler))).Methods("POST")
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

type VerifyOTPResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (uApi *UserAPI) VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	type VerifyOTPRequest struct {
		Email    string `json:"email"`
		OTP      string `json:"otp"`
		DeviceId string `json:"deviceId"`
		Platform string `json:"platform"`
		Location string `json:"location"`
	}
	var req VerifyOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	code, resp, err := uApi.VerifyOTP(r.Context(), req.Email, req.OTP, req.DeviceId, req.Platform, req.Location)
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

func (uApi *UserAPI) GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	code, resp, err := uApi.GetUserProfile(r.Context())
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

func (uApi *UserAPI) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	type UpdateRequest struct {
		FirstName   *string `json:"firstName"`
		LastName    *string `json:"lastName"`
		Email       *string `json:"email"`
		DateOfBirth *string `json:"dateOfBirth"`
	}
	accessToken, ok := r.Context().Value(token.AccessTokenKey).(types.AccessToken)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	code, err := uApi.Update(r.Context(), accessToken.UserId, req.FirstName, req.LastName, req.Email, req.DateOfBirth)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	type UpdateResponse struct{}
	var resp UpdateResponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
