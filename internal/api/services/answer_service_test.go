package services_test

import (
	"api-service/internal/api/services"
	"api-service/internal/models"
	"context"
	"errors"
	"testing"
)

// фейковый репозиторий
type FakeAnswerRepo struct {
	InsertFn        func(ctx context.Context, questionId int64, userId string, text string) error
	GetAnswersFn    func(ctx context.Context, id int64) (*models.Answer, error)
	DeleteAnswersFn func(ctx context.Context, id int64) error
}

func (f *FakeAnswerRepo) InsertAnswerToQuestion(ctx context.Context, questionId int64, userId string, text string) error {
	return f.InsertFn(ctx, questionId, userId, text)
}

func (f *FakeAnswerRepo) GetAnswerById(ctx context.Context, id int64) (*models.Answer, error) {
	return f.GetAnswersFn(ctx, id)
}

func (f *FakeAnswerRepo) DeleteAnswerById(ctx context.Context, id int64) error {
	return f.DeleteAnswersFn(ctx, id)
}

// тесты

func TestCreateAnswer(t *testing.T) {

	tests := []struct {
		name      string
		text      string
		repoError error
		wantErr   error
	}{
		{
			name:      "успех",
			text:      "нормальный ответ",
			repoError: nil,
			wantErr:   nil,
		},
		{
			name:    "текст слишком короткий",
			text:    "hi",
			wantErr: services.ErrAnswerTextLengthViolation,
		},
		{
			name:    "пустой текст",
			text:    "",
			wantErr: services.ErrAnswerTextRequired,
		},
		{
			name:      "репозиторий вернул ошибку",
			text:      "пример ответа",
			repoError: errors.New("db error"),
			wantErr:   errors.New("db error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// создаем фейковый репозиторий
			fake := &FakeAnswerRepo{
				InsertFn: func(ctx context.Context, questionId int64, userId string, text string) error {
					return tc.repoError
				},
			}

			// создаем сервис
			svc := services.NewAnswerService(fake)

			// вызываем сервис
			err := svc.CreateAnswer(context.Background(), 1, "user1", tc.text)

			// проверка ожидаемой ошибки
			if (err != nil && tc.wantErr == nil) ||
				(err == nil && tc.wantErr != nil) {
				t.Fatalf("Ожидалось: %v, получили: %v", tc.wantErr, err)
			}
		})
	}
}
