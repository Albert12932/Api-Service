package questions

import (
	"api-service/internal/api/services"
	"api-service/pkg"
	"net/http"
)

// GetAllQuestionsHandler Получить список всех вопросов
// @Summary      Получить список всех вопросов
// @Tags         Questions
// @Produce      json
// @Success      200  {array}   models.Question
// @Failure      500  {object}  pkg.ErrorResponse   "Внутренняя ошибка"
// @Router       /questions [get]
func GetAllQuestionsHandler(service *services.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// вызываем сервис
		questions, err := service.GetAllQuestions(ctx)
		if err != nil {
			pkg.WriteJSONError(w, http.StatusInternalServerError, err.Error())
			return
		}

		// успех
		pkg.WriteJSON(w, http.StatusOK, questions)
	}
}
