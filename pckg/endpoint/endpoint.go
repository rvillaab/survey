package endpoint

import (
	"context"
	"log"
	ent "survey/pckg/model"
	s "survey/pckg/service"
	t "survey/pckg/transport"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	CreateQuestionEndpoint     endpoint.Endpoint
	GetAllQuestionsEndpoint    endpoint.Endpoint
	UpdateQuestionEndpoint     endpoint.Endpoint
	DeleteQuestionEndpoint     endpoint.Endpoint
	GetQuestionByIdEndpoint    endpoint.Endpoint
	GetQuestionsByUserEndpoint endpoint.Endpoint
	GetAllAnswersEndpoint      endpoint.Endpoint
}

func MakeCreateQuestionEndpoint(svc s.QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ent.Question)
		v, err := svc.CreateQuestion(req)
		if err != nil {

			return t.QuestionResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func MakeGetallQuestionsEndpoint(svc s.QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetAllQuestions()
		if err != nil {
			log.Fatal(err)
			return v, err
		}
		return v, nil
	}
}

func MakeUpdateQuestionEndpoint(svc s.QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(t.QuestionUpdateRequest)
		v, err := svc.UpdateQuestion(req.V, req.S)
		if err != nil {
			return t.QuestionResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func MakeDeleteQuestionEndpoint(svc s.QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(t.QuestionRequest)
		v, err := svc.DeleteQuestion(req.Str)
		if err != nil {
			return t.QuestionResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func MakeGetQuestionByIdEndpoint(svc s.QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(t.QuestionRequest)
		v, err := svc.GetQuestionById(req.Str)
		if err != nil {
			return t.QuestionResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func MakeGetQuestionsByUserEndpoint(svc s.QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(t.QuestionRequest)
		v, err := svc.GetQuestionsByUser(req.Str)
		if err != nil {
			return t.QuestionResponse{Str: "", Err: err.Error()}, nil
		}
		return v, nil
	}
}

func (e Endpoints) GetAllQuestions(ctx context.Context) (string, error) {

	resp, err := e.GetAllQuestionsEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) CreateQuestion(ctx context.Context) (string, error) {

	resp, err := e.CreateQuestionEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) UpdateQuestion(ctx context.Context) (string, error) {

	resp, err := e.UpdateQuestionEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) DeleteQuestion(ctx context.Context) (string, error) {

	resp, err := e.DeleteQuestionEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) GetQuestionById(ctx context.Context) (string, error) {

	resp, err := e.GetQuestionByIdEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) GetQuestionsByUser(ctx context.Context) (string, error) {

	resp, err := e.GetQuestionsByUserEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}

func (e Endpoints) GetAllAnswers(ctx context.Context) (string, error) {

	resp, err := e.GetAllAnswersEndpoint(ctx, nil)
	if err != nil {
		return "", err
	}

	response := resp.(string)
	return response, nil
}
