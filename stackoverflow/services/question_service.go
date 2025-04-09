package services

import (
	"errors"
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

func (s *QuestionService) VoteQuestion(userID, questionID uint, value int) error {
	if value != 1 && value != -1 {
		return errors.New("invalid vote value")
	}

	question, err := s.questionRepo.FindById(questionID)
	if err != nil {
		return err
	}

	if question.UserID == userID {
		return errors.New("cannot vote on your own question")
	}

	question.VoteCount += value
	if err := s.questionRepo.Update(question); err != nil {
		return err
	}

	repDelta := 5
	if value == -1 {
		repDelta = -2
	}
	return s.userRepo.UpdateReputation(question.UserID, repDelta)
}

func (s *QuestionService) Search(query string) ([]models.Question, error) {
	return s.questionRepo.Search(query)
}
