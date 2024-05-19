package service

import (
	"github.com/AZRV17/Skylang/internal/domain"
	"github.com/AZRV17/Skylang/internal/repository"
)

type ExerciseService struct {
	repository repository.Exercises
}

func NewExerciseService(repository repository.Exercises) *ExerciseService {
	return &ExerciseService{
		repository: repository,
	}
}

func (e ExerciseService) GetExerciseByID(id int) (*domain.Exercise, error) {
	return e.repository.GetExerciseByID(id)
}

func (e ExerciseService) GetAllExercises() ([]domain.Exercise, error) {
	return e.repository.GetAllExercises()
}

func (e ExerciseService) CreateExercise(exerciseInput CreateExerciseInput) (*domain.Exercise, error) {
	exercise := domain.Exercise{
		Name:          exerciseInput.Name,
		Description:   exerciseInput.Description,
		FirstVariant:  exerciseInput.FirstVariant,
		SecondVariant: exerciseInput.SecondVariant,
		ThirdVariant:  exerciseInput.ThirdVariant,
		FourthVariant: exerciseInput.FourthVariant,
		CorrectAnswer: exerciseInput.CorrectAnswer,
		Difficulty:    exerciseInput.Difficulty,
		CourseID:      exerciseInput.CourseID,
	}

	return e.repository.CreateExercise(exercise)
}

func (e ExerciseService) UpdateExercise(exerciseInput UpdateExerciseInput) (*domain.Exercise, error) {
	exercise := domain.Exercise{
		ID:            exerciseInput.ID,
		Name:          exerciseInput.Name,
		Description:   exerciseInput.Description,
		CorrectAnswer: exerciseInput.CorrectAnswer,
		FirstVariant:  exerciseInput.FirstVariant,
		SecondVariant: exerciseInput.SecondVariant,
		ThirdVariant:  exerciseInput.ThirdVariant,
		FourthVariant: exerciseInput.FourthVariant,
		Difficulty:    exerciseInput.Difficulty,
		CourseID:      exerciseInput.CourseID,
	}

	return e.repository.UpdateExercise(exercise)
}

func (e ExerciseService) DeleteExercise(id int) error {
	return e.repository.DeleteExercise(id)
}

func (e ExerciseService) GetExercisesByCourseID(id int) ([]domain.Exercise, error) {
	return e.repository.GetExercisesByCourseID(id)
}
