package questions

import (
	"api-service/internal/api/services"
	"api-service/internal/config"
	"api-service/pkg"
	"errors"
	"net/http"
	"strconv"
)

// GetQuestionWithAnswersHandler Получить вопрос и ответы
// @Summary      Получить вопрос и все ответы
// @Tags         Questions
// @Produce      json
// @Param        id   path      int    true  "ID вопроса"
// @Success      200  {object}  models.Question
// @Failure      400  {object}  pkg.ErrorResponse
// @Failure      404  {object}  pkg.ErrorResponse
// @Failure      500  {object}  pkg.ErrorResponse
// @Router       /questions/{id} [get]
func GetQuestionWithAnswersHandler(service *services.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idStr := r.PathValue("id")
		questionId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			pkg.WriteJSONError(w, http.StatusBadRequest, "некорректный id вопроса")
			return
		}

		question, err := service.GetQuestionWithAnswers(ctx, questionId)
		if err != nil {

			// бизнес ошибки
			switch {
			case errors.Is(err, services.ErrQuestionNotFound):
				pkg.WriteJSONError(w, http.StatusNotFound, err.Error())
				return
			}

			// ошибки БД
			switch {
			case errors.Is(err, config.ErrInvalidData),
				errors.Is(err, config.ErrCheckViolation),
				errors.Is(err, config.ErrFieldRequired):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return

			case errors.Is(err, config.ErrNotFound):
				pkg.WriteJSONError(w, http.StatusNotFound, "вопрос не найден")
				return
			}

			// неизвестные ошибки
			pkg.WriteJSONError(w, http.StatusInternalServerError, "internal error")
			return
		}

		// успех
		pkg.WriteJSON(w, http.StatusOK, question)
	}
}
