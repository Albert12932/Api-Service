package repositories

import (
	"api-service/internal/config"
	"api-service/internal/models"
	"context"
	"gorm.io/gorm"
)

type QuestionsRepo interface {
	InsertQuestion(ctx context.Context, text string) error
	GetAllQuestions(ctx context.Context) ([]models.Question, error)
	GetQuestionWithAnswers(ctx context.Context, id int64) (*models.Question, error)
	DeleteQuestionById(ctx context.Context, id int64) error
}

type QuestionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) *QuestionsRepository {
	return &QuestionsRepository{db: db}
}

func (r *QuestionsRepository) InsertQuestion(ctx context.Context, text string) error {

	question := models.Question{
		Text: text,
	}

	// создаём новую запись в таблице questions
	if err := r.db.WithContext(ctx).Create(&question).Error; err != nil {
		return config.Map(err)
	}

	return nil
}

func (r *QuestionsRepository) GetAllQuestions(ctx context.Context) ([]models.Question, error) {

	var questions []models.Question

	// получаем все строки из таблицы questions
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&questions).Error
	if err != nil {
		return nil, config.Map(err)
	}

	return questions, nil
}

func (r *QuestionsRepository) GetQuestionWithAnswers(ctx context.Context, id int64) (*models.Question, error) {

	var question models.Question

	// ищем вопрос по id и загружаем ответы
	err := r.db.WithContext(ctx).Preload("Answers").First(&question, id).Error
	if err != nil {
		return nil, config.Map(err)
	}

	return &question, nil
}

func (r *QuestionsRepository) DeleteQuestionById(ctx context.Context, questionId int64) error {

	err := r.db.WithContext(ctx).Delete(&models.Question{}, questionId).Error
	if err != nil {
		return config.Map(err)
	}

	return nil
}
