package controllers

import (
	"net/http"
	"stackoverflow/services"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	questionService *services.QuestionService
}

func NewQuestionController(questionService *services.QuestionService) *QuestionController {
	return &QuestionController{questionService: questionService}
}

func (h *QuestionController) PostQuestion(c *gin.Context) {
	userID := c.GetUint("userID")

	var input struct {
		Title string   `json:"title" binding:"required"`
		Body  string   `json:"body" binding:"required"`
		Tags  []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question, err := h.questionService.PostQuestion(userID, input.Title, input.Body, input.Tags)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, question)
}
