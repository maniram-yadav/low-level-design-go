package repositories

import (
	"errors"
	"lld/stackoverflow/models"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
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

func (r *TagRepository) List(page, limit int) ([]models.Tag, error) {
	var tags []models.Tag
	offset := (page - 1) * limit
	err := r.db.Preload("User").Offset(offset).Limit(limit).Order("created_at DESC").Find(&tags).Error
	return tags, err
}

// FindOrCreateByName finds a tag by name or creates it if it doesn't exist
func (r *TagRepository) FindOrCreateByName(name string) (*models.Tag, error) {
	if name == "" {
		return nil, errors.New("tag name cannot be empty")
	}

	var tag models.Tag
	err := r.db.Where("name = ?", name).First(&tag).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			tag = models.Tag{Name: name}
			if err := r.db.Create(&tag).Error; err != nil {
				return nil, err
			}
			return &tag, nil
		}
		return nil, err
	}

	return &tag, nil
}

func (r *TagRepository) GetPopularTags(limit int) ([]models.Tag, error) {
	if limit <= 0 {
		limit = 10 // default limit
	}

	var tags []models.Tag
	err := r.db.Model(&models.Tag{}).
		Select("tags.*, COUNT(question_tags.question_id) as question_count").
		Joins("LEFT JOIN question_tags ON question_tags.tag_id = tags.id").
		Group("tags.id").
		Order("question_count DESC").
		Limit(limit).
		Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *TagRepository) GetTagsForQuestion(questionID uint) ([]models.Tag, error) {
	var tags []models.Tag
	err := r.db.Model(&models.Tag{}).
		Joins("JOIN question_tags ON question_tags.tag_id = tags.id").
		Where("question_tags.question_id = ?", questionID).
		Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *TagRepository) GetQuestionsByTag(tagName string, page, limit int) ([]models.Question, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var questions []models.Question
	err := r.db.Model(&models.Question{}).
		Select("DISTINCT questions.*").
		Joins("JOIN question_tags ON question_tags.question_id = questions.id").
		Joins("JOIN tags ON tags.id = question_tags.tag_id").
		Where("tags.name = ?", tagName).
		Order("questions.created_at DESC").
		Offset(offset).
		Limit(limit).
		Preload("User").
		Find(&questions).Error

	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *TagRepository) UpdateTag(tag *models.Tag) error {
	if tag.Name == "" {
		return errors.New("tag name cannot be empty")
	}

	return r.db.Save(tag).Error
}

func (r *TagRepository) DeleteTag(tagID uint) error {
	// Start a transaction
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete tag associations first
	if err := tx.Where("tag_id = ?", tagID).Delete(&models.Tag{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Then delete the tag
	if err := tx.Delete(&models.Tag{}, tagID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
