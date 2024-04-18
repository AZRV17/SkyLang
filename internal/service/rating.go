package service

import (
	"errors"
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
)

type RatingService struct {
	repo repository.Ratings
}

func NewRatingService(repo repository.Ratings) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) CreateRating(input *CreateRatingInput) (*domain.Rating, error) {
	rating := &domain.Rating{
		CourseID: input.CourseID,
		UserID:   input.UserID,
		Grate:    input.Grate,
	}

	ratings, err := s.repo.GetRatingsByCourseID(input.CourseID)
	if err != nil {
		return nil, err
	}
	for _, r := range ratings {
		if r.UserID == input.UserID {
			return nil, errors.New("user already rated this course")
		}
	}

	return s.repo.CreateRating(*rating)
}

func (s *RatingService) GetRatingsByCourseID(id int) ([]domain.Rating, error) {
	return s.repo.GetRatingsByCourseID(id)
}
