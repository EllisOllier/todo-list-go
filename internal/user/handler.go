package user

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UserResponse struct {
	Username     string `json:"username"`
	UserId       int    `json:"user_id"`
	SessionToken string `json:"session_token"`
}

// CreateAccount godoc
// @Summary Creates a new account for a user
// @Description Creates a new account for a user and returns the username, id and session_token
// @Accept json
// @Produce json
// @Param user body UserRequest true "New account details"
// @Success 201 {object} UserRequest
// @Failure 400 {string} string "Bad Request"
// @Router /user [post]
func (s *UserService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req UserRequest

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
	if req.Password == nil {
		http.Error(w, "Missing password field in body", http.StatusBadRequest)
		return
	}
	newUser := User{Username: req.Username, PasswordHash: req.Password}
	userId, err := s.userRepository.CreateAccount(newUser)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	newUser.ID = *userId

	jwt, err := s.GenerateToken(*userId)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	res := UserResponse{Username: *req.Username, UserId: *userId, SessionToken: jwt}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (s *UserService) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req UserRequest

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
	if req.Password == nil {
		http.Error(w, "Missing password field in body", http.StatusBadRequest)
		return
	}

	givenUser := User{Username: req.Username, PasswordHash: req.Password}
	userId, err := s.userRepository.Login(givenUser)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			http.Error(w, "Incorrect Password", http.StatusBadRequest)
			return
		}
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}

	jwt, err := s.GenerateToken(*userId)
	if err != nil {
		http.Error(w, "Server Error: 500", http.StatusInternalServerError)
		return
	}
	res := UserResponse{Username: *req.Username, UserId: *userId, SessionToken: jwt}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
