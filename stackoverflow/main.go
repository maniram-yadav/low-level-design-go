package main

import (
	"lld/stackoverflow/controllers"
	"lld/stackoverflow/middleware"
	"lld/stackoverflow/models"
	"lld/stackoverflow/repositories"
	"lld/stackoverflow/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/openfga/openfga/pkg/storage/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("stackoverflow.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	if err := db.AutoMigrate(
		&models.User{},
		&models.Badge{},
		&models.Question{},
		&models.Answer{},
		&models.Comment{},
		&models.Tag{},
		&models.Vote{},
	); err != nil {
		log.Fatal("failed to migrate database")
	}

	userRepo := repositories.NewUserRepository(db)
	questionRepo := repositories.NewQuestionRepository(db)
	tagRepo := repositories.NewTagRepository(db)
	answerRepo := repositories.NewAnswerRepository(db)

	userService := services.NewUserService(userRepo)
	questionService := services.NewQuestionService(questionRepo, tagRepo, userRepo)
	answerService := services.NewAnswerService(answerRepo, userRepo, questionRepo)

	r := gin.Default()

	ansController := &controllers.AnswerController{answerService}
	authController := &controllers.AuthController{answerService}
	questionController := &controllers.QuestionController{questionService}
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	authGroup := r.Group("/")

	authGroup.Use(middleware.AuthMiddleware(userService))
	{
		authGroup.POST("/questions", questionController.PostQuestion)
		authGroup.POST("/questions/:id/vote", questionController.VoteQuestion)
		authGroup.POST("/answers", ansController.PostAnswer)
	}

	r.GET("/questions", questionController.ListQuestions)
	r.GET("/questions/:id", questionController.GetQuestion)
	r.GET("/search", questionController.SearchQuestions)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server")
	}
}
