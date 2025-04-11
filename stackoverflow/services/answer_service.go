package services

import (
	"errors"
	"lld/stackoverflow/models"
	"lld/stackoverflow/repositories"
)

type AnswerService struct {
	answerRepo   *repositories.AnswerRepository
	userRepo     *repositories.UserRepository
	questionRepo *repositories.QuestionRepository
}

func NewAnswerService(
	answerRepo *repositories.AnswerRepository,
	userRepo *repositories.UserRepository,
	questionRepo *repositories.QuestionRepository,
) *AnswerService {
	return &AnswerService{
		answerRepo:   answerRepo,
		userRepo:     userRepo,
		questionRepo: questionRepo,
	}
}

func (s *AnswerService) PostAnswer(userID, questionID uint, body string) (*models.Answer, error) {
	if body == "" || len(body) < 30 {
		return nil, errors.New("answer body must be at least 30 characters")
	}

	if _, err := s.questionRepo.FindById(questionID); err != nil {
		return nil, errors.New("question not found")
	}

	answer := &models.Answer{
		Body:       body,
		UserID:     userID,
		QuestionID: questionID,
	}

	if err := s.answerRepo.Create(answer); err != nil {
		return nil, err
	}

	return answer, nil
}

func (s *AnswerService) VoteAnswer(userID, answerID uint, value int) error {
	if value != 1 && value != -1 {
		return errors.New("invalid vote value")
	}

	answer, err := s.answerRepo.FindByID(answerID)
	if err != nil {
		return err
	}

	if answer.UserID == userID {
		return errors.New("cannot vote on your own answer")
	}

	if err := s.answerRepo.UpdateVoteCount(answerID, value); err != nil {
		return err
	}

	repDelta := 10
	if value == -1 {
		repDelta = -2
	}
	return s.userRepo.UpdateReputation(answer.UserID, repDelta)
}

func (s *AnswerService) AcceptAnswer(userID, answerID uint) error {
	answer, err := s.answerRepo.FindByID(answerID)
	if err != nil {
		return err
	}

	question, err := s.questionRepo.FindById(answer.QuestionID)
	if err != nil {
		return err
	}

	if question.UserID != userID {
		return errors.New("only the question author can accept answers")
	}

	if err := s.answerRepo.UnacceptAllForQuestion(question.ID); err != nil {
		return err
	}

	return s.answerRepo.AcceptAnswer(answerID)
}

func (s *AnswerService) GetAnswersForQuestion(questionID uint) ([]models.Answer, error) {
	return s.answerRepo.FindByQuestionID(questionID)
}
