package service

import (
	"errors"
	"fmt"
	ent "survey/pckg/entity"
	"time"
)

type QuestionService interface {
	CreateQuestion(ent.Question) (ent.Question, error)
	GetAllQuestions() ([]ent.Question, error)
	Count() int
	UpdateQuestion(string, ent.Question) (interface{}, error)
	DeleteQuestion(string) (string, error)
	GetQuestionById(string) (ent.Question, error)
	GetQuestionsByUser(string) ([]ent.Question, error)
}

// stringService is a concrete implementation of StringService
type QuestionServiceImpl struct{}

type allQuestions []ent.Question

var Questions = allQuestions{
	{
		ID:          "1",
		Content:     "Qué hora es?",
		Description: "pregunta breve",
		CreatedAt:   time.Now(),
		UserCreated: "Rafa",
		UpdatedAt:   time.Now(),
		UserUpdated: "",
	},
}

func NewQuestionService() QuestionService {
	return QuestionServiceImpl{}
}

func (QuestionServiceImpl) CreateQuestion(question ent.Question) (ent.Question, error) {

	Questions = append(Questions, question)
	return question, nil
}

func (QuestionServiceImpl) Count() int {
	return len(Questions)

}

func (QuestionServiceImpl) GetAllQuestions() ([]ent.Question, error) {

	return Questions, nil
}

func (QuestionServiceImpl) UpdateQuestion(questionId string, updatedQuestion ent.Question) (interface{}, error) {

	for i, singleQuestion := range Questions {
		if singleQuestion.ID == questionId {
			singleQuestion.Content = updatedQuestion.Content
			singleQuestion.Description = updatedQuestion.Description
			singleQuestion.UpdatedAt = time.Now()
			singleQuestion.UserCreated = updatedQuestion.UserCreated
			Questions = append(Questions[:i], singleQuestion)
			return updatedQuestion, nil
		}
	}

	return nil, errors.New("Question not found")

}

func (QuestionServiceImpl) DeleteQuestion(questionId string) (string, error) {

	for i, singleQuestion := range Questions {
		if singleQuestion.ID == questionId {
			Questions = append(Questions[:i], Questions[i+1:]...)
			return fmt.Sprintf("The Question with ID %v has been deleted successfully", questionId), nil
		}
	}

	return "", errors.New("Question not found")
}

func (QuestionServiceImpl) GetQuestionById(questionId string) (ent.Question, error) {

	for _, singleQuestion := range Questions {
		if singleQuestion.ID == questionId {
			return singleQuestion, nil
		}
	}

	return ent.Question{}, errors.New("Question not found")
}

func (QuestionServiceImpl) GetQuestionsByUser(user string) ([]ent.Question, error) {

	var questionsUser []ent.Question

	for i, singleQuestion := range Questions {
		if singleQuestion.UserCreated == user {
			questionsUser = append(questionsUser[:i], singleQuestion)
		}
	}

	if len(questionsUser) > 0 {
		return questionsUser, nil
	}

	return nil, errors.New("Question not found")
}
