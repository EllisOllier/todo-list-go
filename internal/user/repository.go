package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(givenDb *sql.DB) *UserRepository {
	return &UserRepository{
		db: givenDb,
	}
}

func (r *UserRepository) CreateAccount(user User) (int, error) {
	var user_id int
	err := r.db.QueryRow("INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id", user.Username, user.PasswordHash).Scan(&user_id)
	if err != nil {
		return user_id, err
	}
	return user_id, nil
}

// Code from [https://gowebexamples.com/password-hashing/]

// used to hash a unhashed password using bcrypt
func HashPassword(password string) (string, error) {
	// []byte(password) converts password string to a slice of bytes as strings are immutable
	// 14 represents the number of hashing iterations which makes brute force harder/longer
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// used to compare a given password to a hashed password
// and return if that passwords are the same
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // err is false if bcrypt returns ErrMissmatchedHashAndPassword
}
