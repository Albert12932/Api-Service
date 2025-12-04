package routes

import (
	"api-service/internal/api/controllers"
	"api-service/internal/api/repositories"
	"api-service/internal/api/services"
	"gorm.io/gorm"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db *gorm.DB) {
	questionRepo := repositories.NewQuestionsRepository(db)
	createQuestionService := services.NewCreateQuestionService(questionRepo)
	mux.Handle("POST /questions", controllers.CreateQuestionHandler(createQuestionService))
}
