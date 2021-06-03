package mocks

import (
	"survey/pckg/model"
	_ "survey/pckg/service"
)

type GetAllQuestionsFunc func() ([]model.Question, error)

type CreateQuestionFunc func(model.Question) (model.Question, error)
type UpdateQuestionFunc func(string, model.Question) (interface{}, error)
type DeleteQuestionFunc func(string) (string, error)
type GetQuestionByIdFunc func(string) (model.Question, error)
type GetQuestionsByUserFunc func(string) ([]model.Question, error)

type MockQuestionService struct {
	MockGetAllQuestionsFunc    GetAllQuestionsFunc
	MockCreateQuestionFunc     CreateQuestionFunc
	MockUpdateQuestionFunc     UpdateQuestionFunc
	MockDeleteQuestionFunc     DeleteQuestionFunc
	MockGetQuestionByIdFunc    GetQuestionByIdFunc
	MockGetQuestionsByUserFunc GetQuestionsByUserFunc
}

func (m *MockQuestionService) SetGetAllFunc(f GetAllQuestionsFunc) {
	m.MockGetAllQuestionsFunc = f
}

func (m *MockQuestionService) GetAllQuestions() ([]model.Question, error) {
	return m.MockGetAllQuestionsFunc()
}

func (m *MockQuestionService) CreateQuestion(q model.Question) (model.Question, error) {
	return m.MockCreateQuestionFunc(q)
}

func (m *MockQuestionService) UpdateQuestion(s string, q model.Question) (interface{}, error) {
	return m.MockUpdateQuestionFunc(s, q)
}

func (m *MockQuestionService) DeleteQuestion(s string) (string, error) {
	return m.MockDeleteQuestionFunc(s)
}

func (m *MockQuestionService) GetQuestionById(s string) (model.Question, error) {
	return m.MockGetQuestionByIdFunc(s)
}

func (m *MockQuestionService) GetQuestionsByUser(s string) ([]model.Question, error) {
	return m.MockGetQuestionsByUserFunc(s)
}

func NewMockService() *MockQuestionService {
	return &MockQuestionService{}
}
