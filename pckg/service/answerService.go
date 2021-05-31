package service

import (
	ent "survey/pckg/entity"
	"time"
)

type AnswerService interface {
	/* 	CreateQuestion(ent.Question) (ent.Question, error) */
	GetAllAnswers() ([]ent.Answer, error)
	/* 	Count() int
	   	UpdateQuestion(string, ent.Question) (interface{}, error)
	   	DeleteQuestion(string) (string, error)
	   	GetQuestionById(string) (string, error)
	   	GetQuestionsByUser(string) (string, error) */
}

// stringService is a concrete implementation of StringService
type AnswerServiceImpl struct{}

type allAnswers []ent.Answer

var Answers = allAnswers{
	{
		ID:          "1",
		Content:     " las 7:20",
		QuestionId:  "1",
		CreatedAt:   time.Now(),
		UserCreated: "Rafa",
	},
}

func NewAnswerService() AnswerService {
	return AnswerServiceImpl{}
}

/*
func (prod QuestionServiceImpl) CreateQuestion(product ent.Question) (ent.Question, error) {

	Questions = append(Questions, product)
	return product, nil
}

func (QuestionServiceImpl) Count() int {
	return len(Questions)

} */

func (AnswerServiceImpl) GetAllAnswers() ([]ent.Answer, error) {

	return Answers, nil
}

/* func (QuestionServiceImpl) UpdateQuestion(questionId string, updatedQuestion ent.Question) (interface{}, error) {

	for i, singleQuestion := range Questions {
		if singleQuestion.ID == questionId {
			singleQuestion.Content = updatedQuestion.Content
			singleQuestion.Description = updatedQuestion.Description
			singleQuestion.UpdatedAt = time.Now()
			Questions = append(Questions[:i], singleQuestion)
			return updatedQuestion, nil
		}
	}

	return nil, errors.New("Product not found")

}

func (QuestionServiceImpl) DeleteQuestion(questionId string) (string, error) {

	for i, singleQuestion := range Questions {
		if singleQuestion.ID == questionId {
			Questions = append(Questions[:i], Questions[i+1:]...)
			return fmt.Sprintf("The product with ID %v has been deleted successfully", questionId), nil
		}
	}

	return "", errors.New("Product not found")
}

func (QuestionServiceImpl) GetQuestionById(questionId string) (string, error) {

	for i, singleQuestion := range Questions {
		if singleQuestion.ID == questionId {
			Questions = append(Questions[:i], Questions[i+1:]...)
			return fmt.Sprintf("The product with ID %v has been deleted successfully", questionId), nil
		}
	}

	return "", errors.New("Product not found")
}

func (QuestionServiceImpl) GetQuestionsByUser(questionId string) (string, error) {

	for i, singleQuestion := range Questions {
		if singleQuestion.ID == questionId {
			Questions = append(Questions[:i], Questions[i+1:]...)
			return fmt.Sprintf("The product with ID %v has been deleted successfully", questionId), nil
		}
	}

	return "", errors.New("Product not found")
}
*/
