package service

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
)

type CourseService struct {
	repositoryCourse repository.Courses
	repositoryUser   repository.Users
}

func NewCourseService(repository repository.Courses, repositoryUser repository.Users) *CourseService {
	return &CourseService{
		repositoryCourse: repository,
		repositoryUser:   repositoryUser,
	}
}

func (c CourseService) GetCourseByID(id int) (*GetCourseOutput, error) {
	course, err := c.repositoryCourse.GetCourseByID(id)
	if err != nil {
		return nil, err
	}

	return &GetCourseOutput{
		Course: *course,
		Author: *course.Author,
	}, nil
}

func (c CourseService) GetAllCourses() ([]GetCourseOutput, error) {
	var output []GetCourseOutput

	courses, err := c.repositoryCourse.GetAllCourses()
	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		output = append(output, GetCourseOutput{
			Course: course,
			Author: *course.Author,
		})
	}

	return output, nil
}

func (c CourseService) CreateCourse(courseInput CreateCourseInput) (*domain.Course, error) {
	course := domain.Course{
		Name:        courseInput.Name,
		Description: courseInput.Description,
		Language:    courseInput.Language,
		Icon:        courseInput.Icon,
		Grate:       courseInput.Grate,
		AuthorID:    courseInput.Author,
	}

	return c.repositoryCourse.CreateCourse(course)
}

func (c CourseService) UpdateCourse(courseInput UpdateCourseInput) (*domain.Course, error) {
	course := domain.Course{
		ID:          courseInput.ID,
		Name:        courseInput.Name,
		Description: courseInput.Description,
		Language:    courseInput.Language,
		Icon:        courseInput.Icon,
		Grate:       courseInput.Grate,
		AuthorID:    courseInput.Author,
		Lectures:    courseInput.Lectures,
		Exercises:   courseInput.Exercises,
	}

	return c.repositoryCourse.UpdateCourse(course)
}

func (c CourseService) DeleteCourse(id int) error {
	return c.repositoryCourse.DeleteCourse(id)
}

func (c CourseService) FilterCoursesByTitle(title string) ([]domain.Course, error) {
	return c.repositoryCourse.FilterCoursesByTitle(title)
}

func (c CourseService) GetCourseByUserID(id int) ([]domain.Course, error) {
	return c.repositoryCourse.GetCourseByUserID(id)
}

func (c CourseService) GetCourseByTitle(title string) (*GetCourseOutput, error) {
	course, err := c.repositoryCourse.GetCourseByTitle(title)
	if err != nil {
		return nil, err
	}

	return &GetCourseOutput{
		Course: *course,
		Author: *course.Author,
	}, nil
}

func (c CourseService) SortCourseByTitle() ([]domain.Course, error) {
	return c.repositoryCourse.SortCourseByTitle()
}

func (c CourseService) SortCourseByDate() ([]domain.Course, error) {
	return c.repositoryCourse.SortCourseByDate()
}

func (c CourseService) SortCourseByRating() ([]domain.Course, error) {
	return c.repositoryCourse.SortCourseByRating()
}
