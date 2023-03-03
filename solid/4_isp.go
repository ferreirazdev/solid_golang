package solid

// ISP - Interface Segregation Principle
// The Interface Segregation Principle (ISP) is a software design principle
// that encourages the creation of small, focused interfaces that are tailored
// to the needs of individual clients.

// BADCASE
type AuthService interface {
	Login(username, password string) (string, error)
	Logout(token string) error
	ChangePassword(token, oldPassword, newPassword string) error
	ResetPassword(username, email string) error
	Register(username, password, email string) error
}

// Good Case
type LoginService interface {
	Login(username, password string) (string, error)
	Logout(token string) error
}

type PasswordService interface {
	ChangePassword(token, oldPassword, newPassword string) error
	ResetPassword(username, email string) error
}

type RegistrationService interface {
	Register(username, password, email string) error
}

// By doing so, we can allow each client to depend only on the interface
// that it needs, without being forced to implement methods that it doesn’t need.
// This can make our system more modular, testable, and maintainable.
