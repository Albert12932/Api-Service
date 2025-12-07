package repositories

import (
	"api-service/internal/config"
	"api-service/internal/models"
	"context"
	"gorm.io/gorm"
)

type AnswersRepo interface {
	InsertAnswerToQuestion(ctx context.Context, questionId int64, userId string, text string) error
	GetAnswerById(ctx context.Context, id int64) (*models.Answer, error)
	DeleteAnswerById(ctx context.Context, id int64) error
}

type AnswersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) *AnswersRepository {
	return &AnswersRepository{db: db}
}

func (r *AnswersRepository) InsertAnswerToQuestion(ctx context.Context, questionId int64, userId, text string) error {

	answer := models.Answer{QuestionId: questionId, UserId: userId, Text: text}

	// создаём запись в таблице answers
	if err := r.db.WithContext(ctx).Create(&answer).Error; err != nil {
		return config.Map(err)
	}

	return nil
}

func (r *AnswersRepository) GetAnswerById(ctx context.Context, answerId int64) (*models.Answer, error) {

	answer := models.Answer{}

	// ищем строку по ключу
	if err := r.db.WithContext(ctx).First(&answer, answerId).Error; err != nil {
		return nil, config.Map(err)
	}

	return &answer, nil
}

func (r *AnswersRepository) DeleteAnswerById(ctx context.Context, answerId int64) error {

	err := r.db.WithContext(ctx).Delete(&models.Answer{}, answerId).Error
	if err != nil {
		return config.Map(err)
	}

	return nil
}
