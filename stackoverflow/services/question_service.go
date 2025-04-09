package services

import "lld/stackoverflow/repositories"

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
