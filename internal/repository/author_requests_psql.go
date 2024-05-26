package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
	"log"
)

// Репозиторий для работы с таблицей author_request
type AuthorRequestRepository struct {
	db *gorm.DB
}

// Функция создания репозитория
func NewAuthorRequestRepository(db *gorm.DB) *AuthorRequestRepository {
	return &AuthorRequestRepository{
		db: db,
	}
}

// Функция создания записи
func (r *AuthorRequestRepository) CreateAuthorRequest(authorRequest domain.AuthorRequest) (*domain.AuthorRequest, error) {
	tx := r.db.Begin()

	if err := tx.Create(&authorRequest).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &authorRequest, nil
}

// Функция получения записи
func (r *AuthorRequestRepository) GetAuthorRequestByID(id int) (*domain.AuthorRequest, error) {
	var authorRequest domain.AuthorRequest

	tx := r.db.Begin()

	if err := tx.First(&authorRequest, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &authorRequest, nil
}

// Функция получения всех записей
func (r *AuthorRequestRepository) GetAuthorRequests() ([]domain.AuthorRequest, error) {
	var authorRequests []domain.AuthorRequest

	tx := r.db.Begin()

	if err := tx.Preload("User").Find(&authorRequests).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	log.Println(authorRequests)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return authorRequests, nil
}

// Функция удаления записи
func (r *AuthorRequestRepository) DeleteAuthorRequest(id int) error {
	tx := r.db.Begin()

	if err := tx.Delete(&domain.AuthorRequest{}, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция получения записи
func (r *AuthorRequestRepository) GetAuthorRequestByUserID(id int) (*domain.AuthorRequest, error) {
	var authorRequests domain.AuthorRequest

	tx := r.db.Begin()

	if err := tx.Where("user_id = ?", id).Find(&authorRequests).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &authorRequests, nil
}
