package repositories

import (
	"lld/stackoverflow/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(question *models.Question) error {
	return r.db.Create(question).Error
}

func (r *QuestionRepository) FindById(id uint) (*models.Question, error) {
	var question models.Question
	err := r.db.Preload("User").Preload("Tags").First(&question, id).Error
	return &question, err
}

func (r *QuestionRepository) Update(question *models.Question) error {
	return r.db.Save(question).Error
}

func (r *QuestionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Question{}, id).Error
}

func (r *QuestionRepository) List(page, limit int) ([]models.Question, error) {
	var questions []models.Question
	offset := (page - 1) * limit
	err := r.db.Preload("User").Offset(offset).Limit(limit).Order("created_at DESC").Find(&questions).Error
	return questions, err
}

func (r *QuestionRepository) Search(query string) ([]models.Question, error) {

	var questions []models.Question
	err := r.db.Preload("User").
		Where("	title LIKE ? or body LIKE ?", "%"+query+"%", "%"+query+"%").
		Order("vote_count DESC").
		Find(&questions).Error
	return questions, err

}
