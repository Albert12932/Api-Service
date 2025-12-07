package answers

import (
	"api-service/internal/api/services"
	"api-service/internal/config"
	"api-service/pkg"
	"errors"
	"net/http"
	"strconv"
)

// GetAnswerByIdHandler Получить ответ по ID
// @Summary      Получить ответ по ID
// @Tags         Answers
// @Produce      json
// @Param        id   path      int    true  "ID ответа"
// @Success      200  {object}  models.Answer
// @Failure      400  {object}  pkg.ErrorResponse
// @Failure      404  {object}  pkg.ErrorResponse
// @Failure      409  {object}  pkg.ErrorResponse
// @Failure      500  {object}  pkg.ErrorResponse
// @Router       /answers/{id} [get]
func GetAnswerByIdHandler(service *services.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// получаем id из URL
		idStr := r.PathValue("id")
		answerId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			pkg.WriteJSONError(w, http.StatusBadRequest, "некорректный id ответа")
			return
		}

		// получаем ответ из сервиса
		answer, err := service.GetAnswerById(ctx, answerId)
		if err != nil {

			// бизнес ошибки
			switch {
			case errors.Is(err, services.ErrAnswerNotFound):
				pkg.WriteJSONError(w, http.StatusNotFound, err.Error())
				return

			case errors.Is(err, services.ErrAnswerTextRequired),
				errors.Is(err, services.ErrAnswerTextLengthViolation):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			// ошибки БД
			switch {
			case errors.Is(err, config.ErrFieldRequired),
				errors.Is(err, config.ErrInvalidData),
				errors.Is(err, config.ErrCheckViolation):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return

			case errors.Is(err, config.ErrNotFound):
				pkg.WriteJSONError(w, http.StatusNotFound, "ответ не найден")
				return

			case errors.Is(err, config.ErrAlreadyExists):
				pkg.WriteJSONError(w, http.StatusConflict, err.Error())
				return
			}

			// неизвестные ошибки
			pkg.WriteJSONError(w, http.StatusInternalServerError, "internal error")
			return
		}

		// успех
		pkg.WriteJSON(w, http.StatusOK, answer)
	}
}
