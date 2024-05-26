package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
	"log"
)

// Репозиторий для работы с таблицей users
type UserRepository struct {
	db *gorm.DB
}

// Функция создания репозитория
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Функция аутентификации
func (u UserRepository) SignInByLogin(login, password string) (*domain.User, error) {
	var user domain.User

	tx := u.db.Begin()

	if err := tx.Preload("Courses").First(&user, "login = ?", login).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция аутентификации по почте
func (u UserRepository) SignInByEmail(email, password string) (*domain.User, error) {
	var user domain.User

	tx := u.db.Begin()

	if err := tx.Preload("Courses").First(&user, "email = ?", email).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция регистрации
func (u UserRepository) SignUp(user domain.User) (*domain.User, error) {
	tx := u.db.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция получения пользователя
func (u UserRepository) GetUserByID(id int) (*domain.User, error) {
	var user domain.User

	tx := u.db.Begin()

	if err := tx.Preload("Courses").First(&user, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	log.Println(user)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция получения всех пользователей
func (u UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User

	tx := u.db.Begin()

	if err := tx.Preload("Courses").Find(&users).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return users, nil
}

// Функция обновления пользователя
func (u UserRepository) UpdateUser(user domain.User) (*domain.User, error) {
	tx := u.db.Begin()

	if err := tx.Preload("Courses").Where("id = ?", user.ID).Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция обновления логина и почты
func (u UserRepository) UpdateUserLoginAndEmail(id int, login, email string) (*domain.User, error) {
	tx := u.db.Begin()

	if err := tx.Model(&domain.User{}).Where("id = ?", id).Update("login", login).Update("email", email).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return nil, nil
}

// Функция удаления
func (u UserRepository) DeleteUser(id int) error {
	tx := u.db.Begin()

	if err := tx.Delete(&domain.User{}, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция обновления пароля
func (u UserRepository) UpdatePassword(id int, password string) (*domain.User, error) {
	tx := u.db.Begin()

	if err := tx.Model(&domain.User{}).Where("id = ?", id).Update("password", password).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var user domain.User

	if err := tx.First(&user, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция обновления курса пользователя
func (u UserRepository) SignUpForCourse(userID, courseID int) error {
	tx := u.db.Begin()

	if err := tx.Model(&domain.User{}).Where("id = ?", userID).Update("course_id", courseID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция обновления пароля по почте
func (u UserRepository) UpdatePasswordByEmail(email, password string) (*domain.User, error) {
	tx := u.db.Begin()

	if err := tx.Model(&domain.User{}).Where("email = ?", email).Update("password", password).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var user domain.User

	if err := tx.First(&user, "email = ?", email).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция обновления аватара
func (u UserRepository) SetUserAvatar(id int, avatar string) (*domain.User, error) {
	tx := u.db.Begin()

	if err := tx.Model(&domain.User{}).Where("id = ?", id).Update("avatar", avatar).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var user domain.User

	if err := tx.First(&user, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// Функция обновления статуса пользователя
func (u UserRepository) UpdateUserCourseStatus(userID, courseID int, status string) error {
	tx := u.db.Begin()

	if err := tx.Model(&domain.User{}).Where("user_id = ? AND course_id = ?", userID, courseID).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция добавления пользователя в курс
func (u UserRepository) CreateUserCourse(userID, courseID int) error {
	tx := u.db.Begin()

	if err := tx.Model(&domain.UserCourse{}).Create(&domain.UserCourse{
		UserID:   userID,
		CourseID: courseID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция удаления пользователя из курса
func (u UserRepository) RemoveUserCourse(userID, courseID int) error {
	tx := u.db.Begin()

	if err := tx.Delete(&domain.UserCourse{}, "user_id = ? AND course_id = ?", userID, courseID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция обновления роли пользователя
func (u UserRepository) UpdateUserRole(userID int, role string) error {
	tx := u.db.Begin()

	if err := tx.Model(&domain.User{}).Where("id = ?", userID).Update("role", role).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
