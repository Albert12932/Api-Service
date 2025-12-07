package answers

import (
	"api-service/internal/api/services"
	"api-service/internal/config"
	"api-service/pkg"
	"errors"
	"net/http"
	"strconv"
)

// DeleteAnswerHandler Удалить ответ
// @Summary      Удалить ответ
// @Tags         Answers
// @Param        id   path      int  true  "ID ответа"
// @Success      200  {object}  map[string]string  "status: deleted"
// @Failure      400  {object}  pkg.ErrorResponse
// @Failure      404  {object}  pkg.ErrorResponse
// @Failure      500  {object}  pkg.ErrorResponse
// @Router       /answer/{id} [delete]
func DeleteAnswerHandler(service *services.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// получаем id из URL
		idStr := r.PathValue("id")
		answerId, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			pkg.WriteJSONError(w, http.StatusBadRequest, "некорректный id ответа")
			return
		}

		// вызываем сервис
		err = service.DeleteAnswer(ctx, answerId)
		if err != nil {

			// бизнес ошибки
			switch {
			case errors.Is(err, services.ErrAnswerNotFound):
				pkg.WriteJSONError(w, http.StatusNotFound, err.Error())
				return
			}

			// ошибки бд
			switch {
			case errors.Is(err, config.ErrInvalidData),
				errors.Is(err, config.ErrCheckViolation),
				errors.Is(err, config.ErrFieldRequired):
				pkg.WriteJSONError(w, http.StatusBadRequest, err.Error())
				return

			case errors.Is(err, config.ErrNotFound):
				pkg.WriteJSONError(w, http.StatusNotFound, "ответ не найден")
				return
			}

			// неизвестные ошибки
			pkg.WriteJSONError(w, http.StatusInternalServerError, "internal error")
			return
		}

		// успех
		pkg.WriteJSON(w, http.StatusOK, map[string]string{
			"status": "deleted",
		})
	}
}
