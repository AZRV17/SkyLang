package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

// Репозиторий для работы с таблицей lectures
type Lecture struct {
	db *gorm.DB
}

// Функция создания репозитория
func NewLectureRepository(db *gorm.DB) *Lecture {
	return &Lecture{
		db: db,
	}
}

// Функция получения записи
func (l Lecture) GetLectureByID(id int) (*domain.Lecture, error) {
	var lecture domain.Lecture

	tx := l.db.Begin()

	if err := tx.Preload("Course").First(&lecture, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &lecture, nil
}

// Функция получения всех записей
func (l Lecture) GetAllLectures() ([]domain.Lecture, error) {
	var lectures []domain.Lecture

	tx := l.db.Begin()

	if err := tx.Preload("Course").Find(&lectures).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return lectures, nil
}

// Функция создания записи
func (l Lecture) CreateLecture(lecture domain.Lecture) (*domain.Lecture, error) {
	tx := l.db.Begin()

	if err := tx.Create(&lecture).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Preload("Course").Last(&lecture).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &lecture, nil
}

// Функция обновления записи
func (l Lecture) UpdateLecture(lecture domain.Lecture) (*domain.Lecture, error) {
	tx := l.db.Begin()

	if err := tx.Where("id = ?", lecture.ID).Save(&lecture).Error; err != nil {
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

// Функция удаления
func (l Lecture) DeleteLecture(id int) error {
	tx := l.db.Begin()

	if err := tx.Where("id = ?", id).Delete(&domain.Lecture{}, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Функция получения записей по курсу
func (l Lecture) GetLecturesByCourseID(courseID int) ([]domain.Lecture, error) {
	var lectures []domain.Lecture

	tx := l.db.Begin()

	if err := tx.Find(&lectures, "course_id = ?", courseID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return lectures, nil
}
