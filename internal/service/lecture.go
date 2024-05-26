package service

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
)

type LectureService struct {
	repository repository.Lectures
}

func NewLectureService(repository repository.Lectures) *LectureService {
	return &LectureService{
		repository: repository,
	}
}

// Функция для получения лекции по ID
func (l LectureService) GetLectureByID(id int) (*domain.Lecture, error) {
	return l.repository.GetLectureByID(id)
}

// Функция для получения всех лекции
func (l LectureService) GetAllLectures() ([]domain.Lecture, error) {
	return l.repository.GetAllLectures()
}

// Функция для создания лекции
func (l LectureService) CreateLecture(lectureInput CreateLectureInput) (*domain.Lecture, error) {
	lecture := domain.Lecture{
		Name:        lectureInput.Name,
		Description: lectureInput.Description,
		CourseID:    lectureInput.Course,
	}

	return l.repository.CreateLecture(lecture)
}

// Функция для обновления лекции
func (l LectureService) UpdateLecture(lectureInput UpdateLectureInput) (*domain.Lecture, error) {
	lecture := domain.Lecture{
		ID:          lectureInput.ID,
		Name:        lectureInput.Name,
		Description: lectureInput.Description,
		CourseID:    int(lectureInput.CourseID),
	}

	return l.repository.UpdateLecture(lecture)
}

func (l LectureService) DeleteLecture(id int) error {
	return l.repository.DeleteLecture(id)
}

func (l LectureService) GetLecturesByCourseID(courseID int) ([]domain.Lecture, error) {
	return l.repository.GetLecturesByCourseID(courseID)
}
