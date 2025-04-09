package repositories

import (
	"lld/stackoverflow/models"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *TagRepository) Create(tag *models.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagRepository) FindById(id uint) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.Preload("Tag").First(&tag, id).Error
	return &tag, err
}

func (r *TagRepository) Update(tag *models.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagRepository) Delete(id uint) error {
	return r.db.Delete(&models.Tag{}, id).Error
}

func (r *TagRepository) List(page, limit int) ([]models.Tag, error) {
	var tags []models.Tag
	offset := (page - 1) * limit
	err := r.db.Preload("User").Offset(offset).Limit(limit).Order("created_at DESC").Find(&tags).Error
	return tags, err
}
