package repository

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"gorm.io/gorm"
)

type Users interface {
	SignInByLogin(login, password string) (*domain.User, error)
	SignInByEmail(email, password string) (*domain.User, error)
	SignUp(user domain.User) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(user domain.User) (*domain.User, error)
	DeleteUser(id int) error
	UpdatePassword(id int, password string) (*domain.User, error)
	SignUpForCourse(userID, courseID int) error
}

type Courses interface {
	GetCourseByID(id int) (*domain.Course, error)
	GetAllCourses() ([]domain.Course, error)
	CreateCourse(course domain.Course) (*domain.Course, error)
	UpdateCourse(course domain.Course) (*domain.Course, error)
	DeleteCourse(id int) error
	GetCourseByTitle(title string) (*domain.Course, error)
	FilterCoursesByTitle(filter string) ([]domain.Course, error)
	GetCourseByUserID(id int) ([]domain.Course, error)
	SortCourseByTitle() ([]domain.Course, error)
	SortCourseByDate() ([]domain.Course, error)
	SortCourseByRating() ([]domain.Course, error)
}

type Lectures interface {
	GetLectureByID(id int) (*domain.Lecture, error)
	GetAllLectures() ([]domain.Lecture, error)
	CreateLecture(lecture domain.Lecture) (*domain.Lecture, error)
	UpdateLecture(lecture domain.Lecture) (*domain.Lecture, error)
	DeleteLecture(id int) error
}

type Exercises interface {
	GetExerciseByID(id int) (*domain.Exercise, error)
	GetAllExercises() ([]domain.Exercise, error)
	CreateExercise(exercise domain.Exercise) (*domain.Exercise, error)
	UpdateExercise(exercise domain.Exercise) (*domain.Exercise, error)
	DeleteExercise(id int) error
}

type Comments interface {
	GetCommentByID(id int) (*domain.Comment, error)
	GetAllComments() ([]domain.Comment, error)
	CreateComment(comment domain.Comment) (*domain.Comment, error)
	UpdateComment(comment domain.Comment) (*domain.Comment, error)
	DeleteComment(id int) error
	GetCommentsByCourseID(id int) ([]domain.Comment, error)
}

type Repository struct {
	Users     Users
	Courses   Courses
	Lectures  Lectures
	Exercises Exercises
	Comments  Comments
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users:     NewUserRepository(db.Model(domain.User{})),
		Courses:   NewCourseRepository(db.Model(domain.Course{})),
		Lectures:  NewLectureRepository(db.Model(domain.Lecture{})),
		Exercises: NewExerciseRepository(db.Model(domain.Exercise{})),
		Comments:  NewCommentRepository(db.Model(domain.Comment{})),
	}
}
