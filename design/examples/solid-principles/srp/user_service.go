package srp

// BAD: Violates SRP - multiple responsibilities
type UserServiceBad struct{}

func (s *UserServiceBad) CreateUser(email, password string) error {
	// 1. Validates user data
	// 2. Sends welcome email
	// 3. Logs the action
	// 4. Saves to database
	return nil // Multiple responsibilities!
}

// GOOD: Follows SRP - single responsibility
type UserService struct {
	validator      *UserValidator
	emailSender    *EmailSender
	logger         *Logger
	userRepository *UserRepository
}

func (s *UserService) CreateUser(email, password string) error {
	// Delegate to specialized components
	if err := s.validator.Validate(email, password); err != nil {
		return err
	}
	
	user := &User{Email: email, Password: password}
	
	if err := s.userRepository.Save(user); err != nil {
		return err
	}
	
	s.emailSender.SendWelcomeEmail(user)
	s.logger.Log("User created: " + email)
	
	return nil
}

type User struct {
	Email    string
	Password string
}

type UserValidator struct{}

func (v *UserValidator) Validate(email, password string) error {
	// Single responsibility: validation only
	return nil
}

type EmailSender struct{}

func (e *EmailSender) SendWelcomeEmail(user *User) {
	// Single responsibility: sending emails
}

type Logger struct{}

func (l *Logger) Log(message string) {
	// Single responsibility: logging
}

type UserRepository struct{}

func (r *UserRepository) Save(user *User) error {
	// Single responsibility: persistence
	return nil
}
