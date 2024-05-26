package service

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
)

type CommentService struct {
	repository repository.Comments
}

func NewCommentService(repository repository.Comments) *CommentService {
	return &CommentService{repository: repository}
}

// Функция для получения комментария по ID
func (s *CommentService) GetCommentByID(id int) (*domain.Comment, error) {
	return s.repository.GetCommentByID(id)
}

// Функция для получения всех комментариев
func (s *CommentService) GetAllComments() ([]domain.Comment, error) {
	return s.repository.GetAllComments()
}

// Функция для создания комментария
func (s *CommentService) CreateComment(commentInput CreateCommentInput) (*domain.Comment, error) {
	comment := domain.Comment{
		Content:  commentInput.Content,
		CourseID: commentInput.CourseID,
		AuthorID: commentInput.UserID,
	}

	return s.repository.CreateComment(comment)
}

// Функция для обновления комментария
func (s *CommentService) UpdateComment(id int, commentInput UpdateCommentInput) (*domain.Comment, error) {

	comment := domain.Comment{
		ID:       commentInput.ID,
		Content:  commentInput.Content,
		CourseID: commentInput.CourseID,
		AuthorID: commentInput.UserID,
	}

	return s.repository.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(id int) error {
	return s.repository.DeleteComment(id)
}

func (s *CommentService) GetCommentsByCourseID(id int) ([]domain.Comment, error) {
	return s.repository.GetCommentsByCourseID(id)
}
