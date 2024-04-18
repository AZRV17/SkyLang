package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

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

func (u UserRepository) UpdateUser(user domain.User) (*domain.User, error) {
	tx := u.db.Begin()

	if err := tx.Preload("Courses").Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

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
