package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"soccer-api/internal/repo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AuthHandler struct {
	AuthRepo repo.UserRepo
	TeamRepo repo.TeamRepo
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.AuthRepo.GetByEmail(r.Context(), credentials.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	hashFromDB := user.PasswordHash
	userID := user.ID

	if err := bcrypt.CompareHashAndPassword([]byte(hashFromDB), []byte(credentials.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "internal_error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if credentials.Email == "" || len(credentials.Password) < 8 {
		http.Error(w, "Invalid email or password (min 8 chars)", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}

	id, err := h.AuthRepo.Create(r.Context(), credentials.Email, string(hash))
	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	if err := h.TeamRepo.AssignNewTeam(r.Context(), id, credentials.Email); err != nil {
		http.Error(w, "User created but team assignment failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}
