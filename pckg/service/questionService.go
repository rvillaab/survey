package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	ent "survey/pckg/model"

	_ "github.com/lib/pq"
)

type QuestionService interface {
	CreateQuestion(ent.Question) (ent.Question, error)
	GetAllQuestions() ([]ent.Question, error)
	UpdateQuestion(string, ent.Question) (interface{}, error)
	DeleteQuestion(string) (string, error)
	GetQuestionById(string) (ent.Question, error)
	GetQuestionsByUser(string) ([]ent.Question, error)
}

// stringService is a concrete implementation of StringService
type QuestionServiceImpl struct {
	dao ent.QuestionDao
}

type allQuestions []ent.Question

func NewQuestionService(host, user, password, dbname string) QuestionService {

	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	var err error
	dB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("BD:", err)
	}

	ent.EnsureTableExists(dB)

	questionDao := ent.QuestionDao{DB: dB}

	return QuestionServiceImpl{dao: questionDao}
}

func (qstImpl QuestionServiceImpl) CreateQuestion(question ent.Question) (ent.Question, error) {

	qstImpl.dao.CreateQuestion(&question)

	return question, nil
}

func (qstImpl QuestionServiceImpl) GetAllQuestions() ([]ent.Question, error) {

	questions, err := qstImpl.dao.GetQuestions(0, 20)
	if err != nil {
		log.Fatal(err)
	}

	return questions, nil
}

func (qstImpl QuestionServiceImpl) UpdateQuestion(questionId string, updatedQuestion ent.Question) (interface{}, error) {

	currentQuestion, err := qstImpl.GetQuestionById(questionId)

	if err != nil {
		log.Print(err)
		return "", err
	}

	currentQuestion.Content = updatedQuestion.Content
	currentQuestion.Description = updatedQuestion.Description
	currentQuestion.Answer = updatedQuestion.Answer

	errUpd := qstImpl.dao.UpdateQuestion(&currentQuestion)
	if errUpd != nil {
		log.Print("Error in update:", errUpd)
		return "", err
	}

	return updatedQuestion, nil

}

func (qstImpl QuestionServiceImpl) DeleteQuestion(questionId string) (string, error) {

	currentQuestion, err := qstImpl.GetQuestionById(questionId)

	if err != nil {
		log.Print(err)
		return "", err
	}

	errDel := qstImpl.dao.DeleteQuestion(currentQuestion.ID)

	if errDel != nil {
		log.Print(errDel)
		return "", errDel
	}

	return fmt.Sprintf("The Question with ID %v has been deleted successfully", currentQuestion.ID), nil
}

func (qstImpl QuestionServiceImpl) GetQuestionById(questionId string) (ent.Question, error) {

	currentQuestion := ent.Question{ID: questionId}

	err := qstImpl.dao.GetQuestion(&currentQuestion)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return ent.Question{}, errors.New("Question not found")
		}
		log.Print(err)
		return ent.Question{}, err
	}

	return currentQuestion, nil
}

func (qstImpl QuestionServiceImpl) GetQuestionsByUser(user string) ([]ent.Question, error) {

	return nil, errors.New("Question not found")
}
