package solid

import "errors"

// SRP - Single Responsibility Principle

// In this example, we have separated the concerns of user management and
// user authentication into two separate responsibilities. The UserRepository
// interface defines methods for interacting with the database, while the UserService
// contains the business logic for user management and authentication.

type UserRepository interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	SaveUser(user *User) error
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserService struct {
	userRepository UserRepository
}

func (s *UserService) RegisterUser(name, email, password string) error {
	user := &User{Name: name, Email: email, Password: password}
	return s.userRepository.SaveUser(user)
}

func (s *UserService) AuthenticateUser(email, password string) (*User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
