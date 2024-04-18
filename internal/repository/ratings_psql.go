package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *RatingRepository {
	return &RatingRepository{
		db: db,
	}
}

func (r RatingRepository) CreateRating(rating domain.Rating) (*domain.Rating, error) {
	tx := r.db.Begin()

	if err := tx.Create(&rating).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &rating, nil
}

func (r RatingRepository) GetRatingsByCourseID(id int) ([]domain.Rating, error) {
	var ratings []domain.Rating

	tx := r.db.Begin()

	if err := tx.Where("course_id = ?", id).Find(&ratings).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return ratings, nil
}
