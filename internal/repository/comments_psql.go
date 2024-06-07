package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

// Репозиторий для работы с таблицей comments
type CommentRepository struct {
	db *gorm.DB
}

// Функция создания репозитория
func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

// Функция получения записи
func (c CommentRepository) GetCommentByID(id int) (*domain.Comment, error) {
	var comment domain.Comment

	tx := c.db.Begin()

	if err := tx.Preload("Author").First(&comment, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, nil
}

// Функция получения всех записей
func (c CommentRepository) GetAllComments() ([]domain.Comment, error) {
	var comments []domain.Comment

	tx := c.db.Begin()

	if err := tx.Preload("Author").Find(&comments).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comments, nil
}

// Функция создания записи
func (c CommentRepository) CreateComment(comment domain.Comment) (*domain.Comment, error) {
	tx := c.db.Begin()

	if err := tx.Create(&comment).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, nil
}

// Функция обновления
func (c CommentRepository) UpdateComment(comment domain.Comment) (*domain.Comment, error) {
	tx := c.db.Begin()

	if err := tx.Where("id = ?", comment.ID).Save(&comment).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, nil
}

// Функция удаления
func (c CommentRepository) DeleteComment(id int) error {
	tx := c.db.Begin()

	if err := tx.Where("id = ?", id).Delete(id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция получения записей
func (c CommentRepository) GetCommentsByCourseID(id int) ([]domain.Comment, error) {
	var comments []domain.Comment

	tx := c.db.Begin()

	if err := tx.Preload("Author").Where("course_id = ?", id).Find(&comments).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comments, nil
}
