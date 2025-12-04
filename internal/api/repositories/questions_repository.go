package repositories

import (
	"api-service/internal/config"
	"api-service/internal/models"
	"context"
	"gorm.io/gorm"
)

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

	if err := r.db.WithContext(ctx).Create(&question).Error; err != nil {
		return config.Map(err)
	}

	return nil

}
