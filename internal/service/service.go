package service

import (
	"github.com/AZRV17/Skylang/internal/config"
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
	"os"
)

type CreateUserInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role"`
}

type UpdateUserInput struct {
	ID      int             `json:"id"`
	Login   string          `json:"login"`
	Email   string          `json:"email"`
	Role    string          `json:"role"`
	Courses []domain.Course `json:"courses"`
}

type Users interface {
	SignInByLogin(login, password string) (*domain.User, error)
	SignInByEmail(email, password string) (*domain.User, error)
	SignUp(userInput CreateUserInput) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(userInput UpdateUserInput) (*domain.User, error)
	UpdatePassword(id int, password string) (*domain.User, error)
	DeleteUser(id int) error
	SignUpForCourse(userID, courseID int) error
	ResetPassword(email string) (int, error)
	UpdatePasswordByEmail(email, password string) (*domain.User, error)
	CreateUserCourse(userID, courseID int) error
	UpdateUserCourseStatus(userID, courseID int, status string) error
	DeleteUserCourse(userID, courseID int) error
}

type CreateCourseInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Icon        string `json:"icon"`
	Author      int    `json:"author"`
}

type UpdateCourseInput struct {
	ID          int               `json:"ID"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Language    string            `json:"language"`
	Icon        string            `json:"icon"`
	Grate       float32           `json:"grate"`
	ReviewCount int               `json:"review_count"`
	Author      int               `json:"author"`
	Lectures    []domain.Lecture  `json:"lectures"`
	Exercises   []domain.Exercise `json:"exercises"`
}

type GetCourseOutput struct {
	Course domain.Course `json:"course"`
	Author domain.User   `json:"author"`
}

type Courses interface {
	GetCourseByID(id int) (*GetCourseOutput, error)
	GetAllCourses() ([]GetCourseOutput, error)
	CreateCourse(courseInput CreateCourseInput) (*domain.Course, error)
	UpdateCourse(courseInput UpdateCourseInput) (*domain.Course, error)
	DeleteCourse(id int) error
	GetCourseByTitle(title string) (*GetCourseOutput, error)
	FilterCoursesByTitle(filter string) ([]domain.Course, error)
	SortCourseByTitle() ([]domain.Course, error)
	SortCourseByDate() ([]domain.Course, error)
	SortCourseByRating() ([]domain.Course, error)
	UpdateCourseGrate(id int, grate *CreateRatingInput) error
}

type CreateRatingInput struct {
	UserID   int
	CourseID int
	Grate    int
}

type Ratings interface {
	CreateRating(input *CreateRatingInput) (*domain.Rating, error)
	GetRatingsByCourseID(id int) ([]domain.Rating, error)
}

type CreateLectureInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Course      int    `json:"course"`
}

type UpdateLectureInput struct {
	ID          int    `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CourseID    uint   `json:"courseID"`
}

type Lectures interface {
	GetLectureByID(id int) (*domain.Lecture, error)
	GetAllLectures() ([]domain.Lecture, error)
	CreateLecture(lectureInput CreateLectureInput) (*domain.Lecture, error)
	UpdateLecture(lectureInput UpdateLectureInput) (*domain.Lecture, error)
	DeleteLecture(id int) error
	GetLecturesByCourseID(courseID int) ([]domain.Lecture, error)
}

type CreateExerciseInput struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	CorrectAnswer string `json:"correctAnswer"`
	Difficulty    string `json:"difficulty"`
	CourseID      uint   `json:"courseID"`
}

type UpdateExerciseInput struct {
	ID            int    `json:"ID"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CorrectAnswer string `json:"correctAnswer"`
	Difficulty    string `json:"difficulty"`
	CourseID      uint   `json:"courseID"`
}

type Exercises interface {
	GetExerciseByID(id int) (*domain.Exercise, error)
	GetAllExercises() ([]domain.Exercise, error)
	CreateExercise(exerciseInput CreateExerciseInput) (*domain.Exercise, error)
	UpdateExercise(exerciseInput UpdateExerciseInput) (*domain.Exercise, error)
	DeleteExercise(id int) error
}

type CreateCommentInput struct {
	Content  string `json:"content"`
	CourseID int    `json:"course_id"`
	UserID   int    `json:"user_id"`
}

type UpdateCommentInput struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	CourseID int    `json:"course_id"`
	UserID   int    `json:"user_id"`
}

type Comments interface {
	GetCommentByID(id int) (*domain.Comment, error)
	GetAllComments() ([]domain.Comment, error)
	CreateComment(commentInput CreateCommentInput) (*domain.Comment, error)
	UpdateComment(id int, commentInput UpdateCommentInput) (*domain.Comment, error)
	DeleteComment(id int) error
	GetCommentsByCourseID(id int) ([]domain.Comment, error)
}

type Email interface {
	SendMailForPasswordReset(recipient string, resetCode int) error
}

type Image interface {
	SetCourseImage(id int, image string) error
	SetUserAvatar(id int, avatar string) error
	GetUserAvatar(id int) (os.File, error)
	GetCourseIcon(id int) (os.File, error)
}

type Service struct {
	repository      repository.Repository
	UserService     Users
	CourseService   Courses
	RatingService   Ratings
	LectureService  Lectures
	ExerciseService Exercises
	CommentService  Comments
	EmailService    Email
	ImageService    Image
}

func NewService(
	repository repository.Repository,
	config config.Config,
) *Service {
	emailService := NewEmailService(config)
	ratingService := NewRatingService(repository.Ratings)

	return &Service{
		repository:      repository,
		UserService:     NewUserService(repository.Users, *emailService),
		CourseService:   NewCourseService(repository.Courses, repository.Users, *ratingService),
		RatingService:   ratingService,
		LectureService:  NewLectureService(repository.Lectures),
		ExerciseService: NewExerciseService(repository.Exercises),
		CommentService:  NewCommentService(repository.Comments),
		EmailService:    emailService,
		ImageService:    NewImageService(repository.Users, repository.Courses),
	}
}
