package service

import (
	"encoding/base64"
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"math/rand"
)

type UserService struct {
	repository   repository.Users
	emailService EmailService
}

func NewUserService(repository repository.Users, emailService EmailService) *UserService {
	return &UserService{
		repository:   repository,
		emailService: emailService,
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
	user, err := u.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	image := user.Avatar

	fileImage, _ := ioutil.ReadFile(image)

	imageToBase64 := base64.StdEncoding.EncodeToString(fileImage)

	user.Avatar = imageToBase64
	return user, nil
}

func (u UserService) GetAllUsers() ([]domain.User, error) {
	return u.repository.GetAllUsers()
}

func (u UserService) UpdateUser(userInput UpdateUserInput) (*domain.User, error) {
	user := domain.User{
		ID:       userInput.ID,
		Login:    userInput.Login,
		Avatar:   userInput.Avatar,
		Password: userInput.Password,
		Email:    userInput.Email,
		Role:     userInput.Role,
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

func (u UserService) UpdateUserLoginAndEmail(id int, login, email string) (*domain.User, error) {
	return u.repository.UpdateUserLoginAndEmail(id, login, email)
}

func (u UserService) DeleteUser(id int) error {
	return u.repository.DeleteUser(id)
}

func (u UserService) SignUpForCourse(userID, courseID int) error {
	return u.repository.SignUpForCourse(userID, courseID)
}

func (u UserService) ResetPassword(email string) (int, error) {
	resetCode := rand.Intn(10000-1000) + 1000

	log.Println(resetCode)

	if err := u.emailService.SendMailForPasswordReset(email, resetCode); err != nil {
		return 0, err
	}

	return resetCode, nil
}

func (u UserService) UpdatePasswordByEmail(email, password string) (*domain.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return u.repository.UpdatePasswordByEmail(email, string(hashPass))
}

func (u UserService) CreateUserCourse(userID, courseID int) error {
	return u.repository.CreateUserCourse(userID, courseID)
}

func (u UserService) UpdateUserCourseStatus(userID, courseID int, status string) error {
	return u.repository.UpdateUserCourseStatus(userID, courseID, status)
}

func (u UserService) DeleteUserCourse(userID, courseID int) error {
	return u.repository.RemoveUserCourse(userID, courseID)
}

func (u UserService) UpdateUserRole(userID int, role string) error {
	return u.repository.UpdateUserRole(userID, role)
}
