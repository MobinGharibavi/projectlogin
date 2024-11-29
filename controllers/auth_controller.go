// controllers/auth_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"project/services"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// هندلر ثبت‌نام
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := services.RegisterUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ثبت‌نام موفقیت‌آمیز بود"))
}

// هندلر ورود
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	success, err := services.LoginUser(req.Username, req.Password)
	if err != nil || !success {
		http.Error(w, "ورود ناموفق بود", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("ورود موفقیت‌آمیز"))
}
