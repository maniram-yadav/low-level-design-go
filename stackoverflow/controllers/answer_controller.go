package handlers

import (
	"net/http"
	"stackoverflow/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnswerController struct {
	answerService *services.AnswerService
}

func NewAnswerController(answerService *services.AnswerService) *AnswerController {
	return &AnswerController{answerService: answerService}
}

func (h *AnswerController) PostAnswer(c *gin.Context) {
	userID := c.GetUint("userID")

	var input struct {
		QuestionID uint   `json:"questionId" binding:"required"`
		Body       string `json:"body" binding:"required,min=30"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	answer, err := h.answerService.PostAnswer(userID, input.QuestionID, input.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, answer)
}

func (h *AnswerController) VoteAnswer(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid answer ID"})
		return
	}

	var input struct {
		Value int `json:"value" binding:"required,oneof=1 -1"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.answerService.VoteAnswer(userID, uint(id), input.Value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "vote recorded"})
}

func (h *AnswerController) AcceptAnswer(c *gin.Context) {
	userID := c.GetUint("userID")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid answer ID"})
		return
	}

	if err := h.answerService.AcceptAnswer(userID, uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "answer accepted"})
}

func (h *AnswerController) GetAnswersForQuestion(c *gin.Context) {
	questionID, err := strconv.Atoi(c.Param("questionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid question ID"})
		return
	}

	answers, err := h.answerService.GetAnswersForQuestion(uint(questionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answers)
}
