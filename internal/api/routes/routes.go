package routes

import (
	"api-service/internal/api/controllers/answers"
	"api-service/internal/api/controllers/questions"
	_ "api-service/internal/api/docs"
	"api-service/internal/api/repositories"
	"api-service/internal/api/services"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/gorm"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db *gorm.DB) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	questionRepo := repositories.NewQuestionsRepository(db)
	answerRepo := repositories.NewAnswersRepository(db)
	questionService := services.NewQuestionService(questionRepo)
	answerService := services.NewAnswerService(answerRepo)
	mux.Handle("POST /questions", questions.CreateQuestionHandler(questionService))
	mux.Handle("GET /questions", questions.GetAllQuestionsHandler(questionService))
	mux.Handle("POST /questions/{id}/answers", answers.CreateAnswerHandler(answerService))
	mux.Handle("GET /answers/{id}", answers.GetAnswerByIdHandler(answerService))
	mux.Handle("GET /questions/{id}", questions.GetQuestionWithAnswersHandler(questionService))
	mux.Handle("DELETE /questions/{id}", questions.DeleteQuestionHandler(questionService))
	mux.Handle("DELETE /answer/{id}", answers.DeleteAnswerHandler(answerService))
}
