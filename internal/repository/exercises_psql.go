package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

type Exercise struct {
	db *gorm.DB
}

func NewExerciseRepository(db *gorm.DB) *Exercise {
	return &Exercise{
		db: db,
	}
}

func (e Exercise) GetExerciseByID(id int) (*domain.Exercise, error) {
	var exercise domain.Exercise

	tx := e.db.Begin()

	if err := tx.First(&exercise, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &exercise, nil
}

func (e Exercise) GetAllExercises() ([]domain.Exercise, error) {
	var exercises []domain.Exercise

	tx := e.db.Begin()

	if err := tx.Find(&exercises).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return exercises, nil
}

func (e Exercise) CreateExercise(exercise domain.Exercise) (*domain.Exercise, error) {
	tx := e.db.Begin()

	if err := tx.Create(&exercise).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Last(&exercise).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &exercise, nil
}

func (e Exercise) UpdateExercise(exercise domain.Exercise) (*domain.Exercise, error) {
	tx := e.db.Begin()

	if err := tx.Where("id = ?", exercise.ID).Save(&exercise).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.First(&exercise, "id = ?", exercise.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &exercise, nil
}

func (e Exercise) DeleteExercise(id int) error {
	tx := e.db.Begin()

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

func (e Exercise) GetExercisesByCourseID(id int) ([]domain.Exercise, error) {
	var exercises []domain.Exercise

	tx := e.db.Begin()

	if err := tx.Find(&exercises, "course_id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return exercises, nil
}
