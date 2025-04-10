package repositories

import (
	"lld/stackoverflow/models"

	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

func (r *AnswerRepository) Create(answer *models.Answer) error {
	return r.db.Create(answer).Error
}

func (r *AnswerRepository) FindByID(id uint) (*models.Answer, error) {
	var answer models.Answer
	err := r.db.Preload("User").First(&answer, id).Error
	return &answer, err
}

func (r *AnswerRepository) FindByQuestionID(questionID uint) ([]models.Answer, error) {
	var answers []models.Answer
	err := r.db.Preload("User").Where("question_id = ?", questionID).Order("created_at DESC").Find(&answers).Error
	return answers, err
}

func (r *AnswerRepository) UpdateVoteCount(answerID uint, delta int) error {
	return r.db.Model(&models.Answer{}).Where("id = ?", answerID).
		Update("vote_count", gorm.Expr("vote_count + ?", delta)).Error
}

func (r *AnswerRepository) AcceptAnswer(answerID uint) error {
	return r.db.Model(&models.Answer{}).Where("id = ?", answerID).
		Update("is_accepted", true).Error
}

func (r *AnswerRepository) UnacceptAllForQuestion(questionID uint) error {
	return r.db.Model(&models.Answer{}).Where("question_id = ?", questionID).
		Update("is_accepted", false).Error
}
