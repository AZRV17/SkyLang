package service

import (
	"encoding/base64"
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
	"io/ioutil"
	"log"
)

type CourseService struct {
	repositoryCourse repository.Courses
	repositoryUser   repository.Users
	serviceRating    RatingService
}

func NewCourseService(repository repository.Courses, repositoryUser repository.Users, ratingService RatingService) *CourseService {
	return &CourseService{
		repositoryCourse: repository,
		repositoryUser:   repositoryUser,
		serviceRating:    ratingService,
	}
}

// Функция для получения курса по ID
func (c CourseService) GetCourseByID(id int) (*domain.Course, error) {
	course, err := c.repositoryCourse.GetCourseByID(id)
	if err != nil {
		return nil, err
	}

	image := course.Icon

	fileImage, _ := ioutil.ReadFile(image)

	imageToBase64 := base64.StdEncoding.EncodeToString(fileImage)

	course.Icon = imageToBase64

	return course, nil
}

// Функция для получения всех курсов
func (c CourseService) GetAllCourses() ([]domain.Course, error) {
	courses, err := c.repositoryCourse.GetAllCourses()
	if err != nil {
		return nil, err
	}

	return courses, nil
}

// Функция для создания курса
func (c CourseService) CreateCourse(courseInput CreateCourseInput) (*domain.Course, error) {
	course := domain.Course{
		Name:        courseInput.Name,
		Description: courseInput.Description,
		Language:    courseInput.Language,
		Icon:        courseInput.Icon,
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
		AuthorID:    courseInput.Author,
		Lectures:    courseInput.Lectures,
		Exercises:   courseInput.Exercises,
		Rating:      int(courseInput.Grate),
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

func (c CourseService) GetCourseByTitle(title string) (*domain.Course, error) {
	course, err := c.repositoryCourse.GetCourseByTitle(title)
	if err != nil {
		return nil, err
	}

	return course, nil
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

func (c CourseService) UpdateCourseGrate(id int, grate *CreateRatingInput) error {
	course, err := c.repositoryCourse.GetCourseByID(id)
	if err != nil {
		return err
	}

	log.Println(grate)

	returnedGrate, err := c.serviceRating.CreateRating(grate)
	if err != nil {
		return err
	}

	course.Grates = append(course.Grates, *returnedGrate)

	totalRate := 0
	for g := range course.Grates {
		totalRate += course.Grates[g].Grate
	}

	course.Rating = totalRate / len(course.Grates)

	_, err = c.repositoryCourse.UpdateCourse(*course)

	return err
}

func (c CourseService) GetCourseByAuthorID(id int) ([]domain.Course, error) {
	return c.repositoryCourse.GetCourseByAuthorID(id)
}
