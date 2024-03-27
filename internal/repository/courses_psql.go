package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (c CourseRepository) GetCourseByID(id int) (*domain.Course, error) {
	var course domain.Course

	tx := c.db.Begin()

	if err := tx.First(&course, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &course, nil
}

func (c CourseRepository) GetAllCourses() ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}

func (c CourseRepository) CreateCourse(course domain.Course) (*domain.Course, error) {
	tx := c.db.Begin()

	if err := tx.Create(&course).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Last(&course).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &course, nil
}

func (c CourseRepository) UpdateCourse(course domain.Course) (*domain.Course, error) {
	tx := c.db.Begin()

	if err := tx.Save(&course).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.First(&course, "id = ?", course.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &course, nil
}

func (c CourseRepository) DeleteCourse(id int) error {
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

func (c CourseRepository) GetCourseByTitle(title string) (*domain.Course, error) {
	var course domain.Course

	tx := c.db.Begin()

	if err := tx.First(&course, "title = ?", title).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &course, nil
}

func (c CourseRepository) FilterCoursesByTitle(filter string) ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Where("title LIKE ?", "%"+filter+"%").Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}

func (c CourseRepository) GetCourseByUserID(id int) ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Where("user_id = ?", id).Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}

func (c CourseRepository) SortCourseByTitle() ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Order("title").Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}

func (c CourseRepository) SortCourseByDate() ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Order("created_at").Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}

func (c CourseRepository) SortCourseByRating() ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Order("rating").Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}

func (c CourseRepository) SetCourseIcon(id int, icon string) error {
	tx := c.db.Begin()

	if err := tx.Model(&domain.Course{}).Where("id = ?", id).Update("icon", icon).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
