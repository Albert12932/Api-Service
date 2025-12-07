package services

import (
	"api-service/internal/api/repositories"
	"api-service/internal/config"
	"api-service/internal/models"
	"context"
	"errors"
)

var (
	ErrQuestionTextRequired        = errors.New("отсутствует текст вопроса")
	ErrQuestionTextLengthViolation = errors.New("текст вопроса слишком короткий")
	ErrQuestionAlreadyExists       = errors.New("такой вопрос уже существует")
	ErrQuestionNotFound            = errors.New("вопрос не найден")
)

type QuestionService struct {
	questionRepo repositories.QuestionsRepo
}

func NewQuestionService(questionRepo repositories.QuestionsRepo) *QuestionService {
	return &QuestionService{questionRepo: questionRepo}
}

func (s *QuestionService) CreateQuestion(ctx context.Context, text string) error {
	if len(text) < 3 {
		return ErrQuestionTextLengthViolation
	}
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

func (s *QuestionService) GetAllQuestions(ctx context.Context) ([]models.Question, error) {
	questions, err := s.questionRepo.GetAllQuestions(ctx)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (s *QuestionService) GetQuestionWithAnswers(ctx context.Context, id int64) (*models.Question, error) {
	question, err := s.questionRepo.GetQuestionWithAnswers(ctx, id)
	if err != nil {

		if errors.Is(err, config.ErrNotFound) {
			return nil, ErrQuestionNotFound
		}

		return nil, err
	}

	return question, nil
}

func (s *QuestionService) DeleteQuestion(ctx context.Context, questionId int64) error {
	err := s.questionRepo.DeleteQuestionById(ctx, questionId)
	if err != nil {

		if errors.Is(err, config.ErrNotFound) {
			return ErrQuestionNotFound
		}

		return err
	}

	return nil
}
