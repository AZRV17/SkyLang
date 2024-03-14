package service

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.Users
}

func NewUserService(repository repository.Users) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u UserService) SignInByLogin(login, password string) (*domain.User, error) {
	// TODO: add JWT token authorization
	user, err := u.repository.SignInByLogin(login, password)
	if err != nil {
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) SignInByEmail(email, password string) (*domain.User, error) {
	user, err := u.repository.SignInByEmail(email, password)
	if err != nil {
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) SignUp(userInput CreateUserInput) (*domain.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := domain.User{
		Login:    userInput.Login,
		Password: string(hashedPass),
		Email:    userInput.Email,
		Role:     userInput.Role,
	}

	return u.repository.SignUp(user)
}

func (u UserService) GetUserByID(id int) (*domain.User, error) {
	return u.repository.GetUserByID(id)
}

func (u UserService) GetAllUsers() ([]domain.User, error) {
	return u.repository.GetAllUsers()
}

func (u UserService) UpdateUser(userInput UpdateUserInput) (*domain.User, error) {
	user := domain.User{
		ID:    userInput.ID,
		Login: userInput.Login,
		Email: userInput.Email,
		Role:  userInput.Role,
	}

	return u.repository.UpdateUser(user)
}

func (u UserService) UpdatePassword(id int, password string) (*domain.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return u.repository.UpdatePassword(id, string(hashPass))
}

func (u UserService) DeleteUser(id int) error {
	return u.repository.DeleteUser(id)
}

func (u UserService) SignUpForCourse(userID, courseID int) error {
	return u.repository.SignUpForCourse(userID, courseID)
}
