package questions

import (
	"api-service/internal/api/services"
	"api-service/internal/config"
	"api-service/pkg"
	"encoding/json"
	"errors"
	"net/http"
)

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

// CreateQuestionHandler Создать новый вопрос
// @Summary      Создать новый вопрос
// @Description  Создаёт вопрос с указанным текстом
// @Tags         Questions
// @Accept       json
// @Produce      json
// @Param        request  body      questions.CreateQuestionRequest  true  "Текст вопроса"
// @Success      201      {object}  map[string]string               "status: ok"
// @Failure      400      {object}  pkg.ErrorResponse               "Ошибка валидации"
// @Failure      409      {object}  pkg.ErrorResponse               "Такой вопрос уже существует"
// @Failure      500      {object}  pkg.ErrorResponse               "Внутренняя ошибка"
// @Router       /questions [post]
func CreateQuestionHandler(service *services.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// получаем текст вопроса из тела запроса
		var req CreateQuestionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			pkg.WriteJSONError(w, http.StatusBadRequest, "некорректные данные")
			return
		}

		// проверяем длину текста
		if len(req.Text) < 3 {
			pkg.WriteJSONError(w, http.StatusBadRequest, services.ErrQuestionTextLengthViolation.Error())
			return
		}

		// вызываем сервис
		err := service.CreateQuestion(ctx, req.Text)
		if err != nil {

			// бизнес ошибки
			switch {
			case errors.Is(err, services.ErrQuestionTextRequired):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return
			case errors.Is(err, services.ErrQuestionAlreadyExists):
				pkg.WriteJSONError(w, http.StatusConflict, err.Error())
				return
			}

			// ошибки БД
			switch {
			case errors.Is(err, config.ErrFieldRequired),
				errors.Is(err, config.ErrInvalidData),
				errors.Is(err, config.ErrCheckViolation):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			// неизвестная ошибка
			pkg.WriteJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// успех
		pkg.WriteJSON(w, http.StatusCreated, map[string]string{
			"status": "ok",
		})
	}
}
