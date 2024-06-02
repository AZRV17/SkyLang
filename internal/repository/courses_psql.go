package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
	"log"
)

// Репозиторий для работы с таблицей courses
type CourseRepository struct {
	db *gorm.DB
}

// Функция создания репозитория
func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Функция получения записи
func (c CourseRepository) GetCourseByID(id int) (*domain.Course, error) {
	var course domain.Course

	tx := c.db.Begin()

	if err := tx.Preload("Author").First(&course, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &course, nil
}

// Функция получения всех записей
func (c CourseRepository) GetAllCourses() ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Preload("Author").Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	log.Println(courses[0].Author)

	return courses, nil
}

// Функция создания записи
func (c CourseRepository) CreateCourse(course domain.Course) (*domain.Course, error) {
	tx := c.db.Begin()

	if err := tx.Create(&course).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Preload("Author").Last(&course).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &course, nil
}

// Функция обновления
func (c CourseRepository) UpdateCourse(course domain.Course) (*domain.Course, error) {
	tx := c.db.Begin()

	if err := tx.Where("id = ?", course.ID).Save(&course).Error; err != nil {
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

// Функция удаления
func (c CourseRepository) DeleteCourse(id int) error {
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

// Функция получения записи
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

// Функция фильтрации
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

// Функция получения записи
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

// Функция сортировки
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

// Функция сортировки
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

// Функция сортировки
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

// Функция установки иконки
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

// Функция получения записей
func (c CourseRepository) GetCourseByAuthorID(id int) ([]domain.Course, error) {
	var courses []domain.Course

	tx := c.db.Begin()

	if err := tx.Preload("Author").Where("author_id = ?", id).Find(&courses).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}
