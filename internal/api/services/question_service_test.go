package services_test

import (
	"api-service/internal/api/services"
	"api-service/internal/models"
	"context"
	"errors"
	"testing"
)

// фейковый репозиторий
type FakeQuestionRepo struct {
	InsertFn func(ctx context.Context, text string) error
	GetAllFn func(ctx context.Context) ([]models.Question, error)
	GetOneFn func(ctx context.Context, id int64) (*models.Question, error)
	DeleteFn func(ctx context.Context, id int64) error
}

func (f *FakeQuestionRepo) InsertQuestion(ctx context.Context, text string) error {
	return f.InsertFn(ctx, text)
}

func (f *FakeQuestionRepo) GetAllQuestions(ctx context.Context) ([]models.Question, error) {
	return f.GetAllFn(ctx)
}

func (f *FakeQuestionRepo) GetQuestionWithAnswers(ctx context.Context, id int64) (*models.Question, error) {
	return f.GetOneFn(ctx, id)
}

func (f *FakeQuestionRepo) DeleteQuestionById(ctx context.Context, id int64) error {
	return f.DeleteFn(ctx, id)
}

// тесты

func TestCreateQuestion(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		repoError error
		wantErr   error
	}{
		{
			name:      "успех",
			input:     "нормальный вопрос",
			repoError: nil,
			wantErr:   nil,
		},
		{
			name:    "короткий вопрос",
			input:   "hi",
			wantErr: services.ErrQuestionTextLengthViolation,
		},
		{
			name:      "репозиторий вернул ошибку",
			input:     "Пример",
			repoError: errors.New("db error"),
			wantErr:   errors.New("db error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			fake := &FakeQuestionRepo{
				InsertFn: func(ctx context.Context, text string) error {
					return tc.repoError
				},
			}

			svc := services.NewQuestionService(fake)

			err := svc.CreateQuestion(context.Background(), tc.input)

			if (err != nil && tc.wantErr == nil) ||
				(err == nil && tc.wantErr != nil) {
				t.Fatalf("Ожидалось: %v, получили: %v", tc.wantErr, err)
			}
		})
	}
}
