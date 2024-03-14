package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (c CommentRepository) GetCommentByID(id int) (*domain.Comment, error) {
	var comment domain.Comment

	tx := c.db.Begin()

	if err := tx.First(&comment, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, nil
}

func (c CommentRepository) GetAllComments() ([]domain.Comment, error) {
	var comments []domain.Comment

	tx := c.db.Begin()

	if err := tx.Find(&comments).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comments, nil
}

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

func (c CommentRepository) UpdateComment(comment domain.Comment) (*domain.Comment, error) {
	tx := c.db.Begin()

	if err := tx.Save(&comment).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, nil
}

func (c CommentRepository) DeleteComment(id int) error {
	tx := c.db.Begin()

	if err := tx.Delete(id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (c CommentRepository) GetCommentsByCourseID(id int) ([]domain.Comment, error) {
	var comments []domain.Comment

	tx := c.db.Begin()

	if err := tx.Where("course_id = ?", id).Find(&comments).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comments, nil
}
