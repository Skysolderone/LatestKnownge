package main

type User struct {
	Id   int
	Name string
}

type UserRepository interface {
	GetById(id int) (User, error)
	Save(user User) error
}
type UserService struct {
	repo   UserRepository
	email  EmailService
	logger Logger
}

type EmailService interface {
	Send(user User) error
}
type Logger interface {
	Info(message string)
	Error(message string)
}

func (u *UserService) NewUser(repo UserRepository, email EmailService, logger Logger) *UserService {
	return &UserService{repo: repo, email: email, logger: logger}
}

func (u *UserService) CreateUser(user User) error {
	u.logger.Info("Creating user")
	if err := u.repo.Save(user); err != nil {
		u.logger.Error("Error creating user")
		return err
	}
	if err := u.email.Send(user); err != nil {
		u.logger.Error("Error sending email")
		return err
	}
	u.logger.Info("User created successfully")
	return nil
}

func main() {
	// println("Hello, World!")
}
