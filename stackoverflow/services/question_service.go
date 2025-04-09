package services

import (
	"lld/stackoverflow/models"
	"lld/stackoverflow/repositories"
)

type QuestionService struct {
	questionRepo *repositories.QuestionRepository
	tagRepo      *repositories.TagRepository
	userRepo     *repositories.UserRepository
}

func NewQuestionService(questionRepo *repositories.QuestionRepository,
	tagRepo *repositories.TagRepository,
	userRepo *repositories.UserRepository) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
		tagRepo:      tagRepo,
		userRepo:     userRepo,
	}
}

func (s *QuestionService) PostQuestion(userID uint, title, body string, tagNames []string) (*models.Question, error) {
	// Get or create tags
	tags := make([]models.Tag, 0)
	for _, name := range tagNames {
		tag, err := s.tagRepo.FindOrCreateByName(name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, *tag)
	}
	question := &models.Question{
		Title:     title,
		Body:      body,
		UserID:    userID,
		Tags:      tags,
		Status:    "open",
		VoteCount: 0,
		Views:     0,
	}
	if err := s.questionRepo.Create(question); err != nil {
		return nil, err
	}
	return question, nil
}

func (s *QuestionService) GetQuestion(id uint) (*models.Question, error) {
	question, err := s.questionRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	// Increment view count
	question.Views++
	_ = s.questionRepo.Update(question)

	return question, nil
}
