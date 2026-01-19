package user

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type NewUserRequest struct {
		Username     *string `json:"username"`
		PasswordHash *string `json:"password_hash"`
	}
	var req NewUserRequest

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request: 400", http.StatusBadRequest)
		return
	}
	if req.Username == nil {
		http.Error(w, "Missing username field in body", http.StatusBadRequest)
		return

	}
	if req.PasswordHash == nil {
		http.Error(w, "Missing password_hash field in body", http.StatusBadRequest)
		return
	}
	newUser := User{Username: *req.Username, PasswordHash: *req.PasswordHash}
	userId, err := s.userRepository.CreateAccount(newUser)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	newUser.ID = userId

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type GivenUserRequest struct {
		Username     *string `json:"username"`
		PasswordHash *string `json:"password_hash"`
	}
	var req GivenUserRequest

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request: 400", http.StatusBadRequest)
		return
	}
	if req.Username == nil {
		http.Error(w, "Missing username field in body", http.StatusBadRequest)
		return

	}
	if req.PasswordHash == nil {
		http.Error(w, "Missing password_hash field in body", http.StatusBadRequest)
		return
	}

	givenUser := User{Username: *req.Username, PasswordHash: *req.PasswordHash}
	userId, err := s.userRepository.Login(givenUser)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			http.Error(w, "Incorrect Password", http.StatusBadRequest)
			return
		}
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userId)
}
