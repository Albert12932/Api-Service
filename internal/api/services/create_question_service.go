package services

import (
	"api-service/internal/api/repositories"
	"api-service/internal/config"
	"context"
	"errors"
)

var (
	ErrQuestionTextRequired        = errors.New("отсутствует текст вопроса")
	ErrQuestionTextLengthViolation = errors.New("текст вопроса слишком короткий")
	ErrQuestionAlreadyExists       = errors.New("такой вопрос уже существует")
)

type CreateQuestionService struct {
	questionRepo *repositories.QuestionsRepository
}

func NewCreateQuestionService(questionRepo *repositories.QuestionsRepository) *CreateQuestionService {
	return &CreateQuestionService{questionRepo: questionRepo}
}

func (s *CreateQuestionService) CreateQuestion(ctx context.Context, text string) error {
	err := s.questionRepo.InsertQuestion(ctx, text)
	if err != nil {
		if errors.Is(err, config.ErrFieldRequired) {
			return ErrQuestionTextRequired
		}
		if errors.Is(err, config.ErrCheckViolation) {
			return ErrQuestionTextLengthViolation
		}
		if errors.Is(err, config.ErrAlreadyExists) {
			return ErrQuestionAlreadyExists
		}
		return err
	}

	return nil
}
