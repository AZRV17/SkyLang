package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

type Lecture struct {
	db *gorm.DB
}

func NewLectureRepository(db *gorm.DB) *Lecture {
	return &Lecture{
		db: db,
	}
}

func (l Lecture) GetLectureByID(id int) (*domain.Lecture, error) {
	var lecture domain.Lecture

	tx := l.db.Begin()

	if err := tx.First(&lecture, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &lecture, nil
}

func (l Lecture) GetAllLectures() ([]domain.Lecture, error) {
	var lectures []domain.Lecture

	tx := l.db.Begin()

	if err := tx.Find(&lectures).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return lectures, nil
}

func (l Lecture) CreateLecture(lecture domain.Lecture) (*domain.Lecture, error) {
	tx := l.db.Begin()

	if err := tx.Create(&lecture).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Last(&lecture).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &lecture, nil
}

func (l Lecture) UpdateLecture(lecture domain.Lecture) (*domain.Lecture, error) {
	tx := l.db.Begin()

	if err := tx.Save(&lecture).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.First(&lecture, "id = ?", lecture.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &lecture, nil
}

func (l Lecture) DeleteLecture(id int) error {
	tx := l.db.Begin()

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
