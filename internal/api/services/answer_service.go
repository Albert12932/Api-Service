package services

import (
	"api-service/internal/api/repositories"
	"api-service/internal/config"
	"api-service/internal/models"
	"context"
	"errors"
)

var (
	ErrAnswerTextRequired        = errors.New("отсутствует текст ответа")
	ErrAnswerTextLengthViolation = errors.New("текст ответа слишком короткий")
	ErrAnswerNotFound            = errors.New("ответ не найден")
	ErrUserIdRequired            = errors.New("не указан пользователь ответа")
)

type AnswerService struct {
	AnswerRepo repositories.AnswersRepo
}

func NewAnswerService(answerRepo repositories.AnswersRepo) *AnswerService {
	return &AnswerService{AnswerRepo: answerRepo}
}

func (s *AnswerService) CreateAnswer(ctx context.Context, questionId int64, userId, text string) error {
	if len(text) == 0 {
		return ErrAnswerTextRequired
	}
	if len(text) < 3 {
		return ErrAnswerTextLengthViolation
	}
	if userId == "" {
		return ErrUserIdRequired
	}

	return s.AnswerRepo.InsertAnswerToQuestion(ctx, questionId, userId, text)
}

func (s *AnswerService) GetAnswerById(ctx context.Context, answerId int64) (*models.Answer, error) {
	answer, err := s.AnswerRepo.GetAnswerById(ctx, answerId)
	if err != nil {
		if errors.Is(err, config.ErrNotFound) {
			return nil, ErrAnswerNotFound
		}
		return nil, err
	}
	return answer, nil
}

func (s *AnswerService) DeleteAnswer(ctx context.Context, answerId int64) error {
	err := s.AnswerRepo.DeleteAnswerById(ctx, answerId)
	if err != nil {

		if errors.Is(err, config.ErrNotFound) {
			return ErrAnswerNotFound
		}

		return err
	}

	return nil
}
