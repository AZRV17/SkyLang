package service

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
)

type AuthorRequestService struct {
	repository  repository.AuthorRequests
	userService UserService
}

func NewAuthorRequestService(repository repository.AuthorRequests, userService UserService) *AuthorRequestService {
	return &AuthorRequestService{
		repository:  repository,
		userService: userService,
	}
}

// Функции получения запросов на авторство
func (a AuthorRequestService) GetAuthorRequests() ([]domain.AuthorRequest, error) {
	return a.repository.GetAuthorRequests()
}

// Функции получения запроса на авторство по ID
func (a AuthorRequestService) GetAuthorRequestByID(id int) (*domain.AuthorRequest, error) {
	authorRequest, err := a.repository.GetAuthorRequestByID(id)

	if err != nil {
		return nil, err
	}

	return authorRequest, nil
}

// Функция создания запроса на авторство
func (a AuthorRequestService) CreateAuthorRequest(authorRequestInput CreateAuthorRequestInput) (*domain.AuthorRequest, error) {
	user, err := a.userService.GetUserByID(authorRequestInput.UserID)
	if err != nil {
		return nil, err
	}

	authorRequest := domain.AuthorRequest{
		UserID: user.ID,
		User:   user,
	}

	return a.repository.CreateAuthorRequest(authorRequest)
}

func (a AuthorRequestService) DeleteAuthorRequest(id int) error {
	authorRequest, err := a.repository.GetAuthorRequestByID(id)
	if err != nil {
		return err
	}

	if authorRequest == nil {
		return nil
	}

	err = a.userService.UpdateUserRole(authorRequest.UserID, "author")
	if err != nil {
		return err
	}

	return a.repository.DeleteAuthorRequest(id)
}

func (a AuthorRequestService) GetAuthorRequestByUserID(id int) (*domain.AuthorRequest, error) {
	return a.repository.GetAuthorRequestByUserID(id)
}
