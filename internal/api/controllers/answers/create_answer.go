package answers

import (
	"api-service/internal/api/services"
	"api-service/internal/config"
	"api-service/internal/models"
	"api-service/pkg"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// CreateAnswerHandler Добавить ответ к вопросу
// @Summary      Добавить ответ к вопросу
// @Description  Создаёт новый ответ для указанного вопроса
// @Tags         Answers
// @Accept       json
// @Produce      json
// @Param        id       path      int                         true  "ID вопроса"
// @Param        request  body      models.CreateAnswerRequest true "Текст ответа и user_id"
// @Success      201      {object}  map[string]string           "status: ok"
// @Failure      400      {object}  pkg.ErrorResponse           "Ошибка валидации"
// @Failure      404      {object}  pkg.ErrorResponse           "Вопрос не найден"
// @Failure      409      {object}  pkg.ErrorResponse           "Ответ уже существует"
// @Failure      500      {object}  pkg.ErrorResponse           "Внутренняя ошибка"
// @Router       /questions/{id}/answers [post]
func CreateAnswerHandler(service *services.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idStr := r.PathValue("id")
		questionId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			pkg.WriteJSONError(w, http.StatusBadRequest, "некорректный id вопроса")
			return
		}

		var req models.CreateAnswerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			pkg.WriteJSONError(w, http.StatusBadRequest, "некорректные данные")
			return
		}

		err = service.CreateAnswer(ctx, questionId, req.UserId, req.Text)
		if err != nil {
			//бизнес ошибки
			switch {
			case errors.Is(err, services.ErrAnswerTextRequired),
				errors.Is(err, services.ErrAnswerTextLengthViolation),
				errors.Is(err, services.ErrUserIdRequired):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			//ошибки БД
			switch {
			case errors.Is(err, config.ErrFieldRequired),
				errors.Is(err, config.ErrInvalidData),
				errors.Is(err, config.ErrCheckViolation):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return

			case errors.Is(err, config.ErrAlreadyExists):
				pkg.WriteJSONError(w, http.StatusConflict, err.Error())
				return
			}

			//неизвестная ошибка
			pkg.WriteJSONError(w, http.StatusInternalServerError, "internal error")
			return
		}

		// успех
		pkg.WriteJSON(w, http.StatusCreated, map[string]string{
			"status": "ok",
		})
	}
}
