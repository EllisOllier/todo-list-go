package user

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"` // stops the return of password_hash
}

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(givenUserRepository *UserRepository) *UserService {
	return &UserService{
		userRepository: givenUserRepository,
	}
}
