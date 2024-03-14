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

func (s *CommentService) GetCommentByID(id int) (*domain.Comment, error) {
	return s.repository.GetCommentByID(id)
}

func (s *CommentService) GetAllComments() ([]domain.Comment, error) {
	return s.repository.GetAllComments()
}

func (s *CommentService) CreateComment(commentInput CreateCommentInput) (*domain.Comment, error) {
	comment := domain.Comment{
		Content:  commentInput.Content,
		CourseID: commentInput.CourseID,
		AuthorID: commentInput.UserID,
	}

	return s.repository.CreateComment(comment)
}

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
