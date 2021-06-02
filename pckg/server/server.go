package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	e "survey/pckg/endpoint"
	ent "survey/pckg/model"
	quest "survey/pckg/question"
	service "survey/pckg/service"
	t "survey/pckg/transport"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints e.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/questions").Handler(httptransport.NewServer(
		endpoints.GetAllQuestionsEndpoint,
		t.DecodeCountRequest,
		t.EncodeResponse,
	))

	r.Methods("DELETE").Path("/questions/{id}").Handler(httptransport.NewServer(
		endpoints.DeleteQuestionEndpoint,
		t.DecodeDeleteRequest,
		t.EncodeResponse,
	))

	r.Methods("GET").Path("/questions/{id}").Handler(httptransport.NewServer(
		endpoints.GetQuestionByIdEndpoint,
		t.DecodeDeleteRequest,
		t.EncodeResponse,
	))

	r.Methods("GET").Path("/questions/user/{name}").Handler(httptransport.NewServer(
		endpoints.GetQuestionsByUserEndpoint,
		t.DecodeGetByUserRequest,
		t.EncodeResponse,
	))

	r.Methods("PUT").Path("/questions/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateQuestionEndpoint,
		t.DecodeUpdateRequest,
		t.EncodeResponse,
	))

	r.Methods("POST").Path("/question").Handler(httptransport.NewServer(
		endpoints.CreateQuestionEndpoint,
		t.DecodeQuestionCreateRequest,
		t.EncodeResponse,
	))

	r.Methods("GET").Path("/answers").Handler(httptransport.NewServer(
		endpoints.GetAllAnswersEndpoint,
		t.DecodeCountRequest,
		t.EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type QuestionGRPCServer struct {
	Serv service.QuestionService
	quest.UnimplementedQuestionServiceServer
}

func (s *QuestionGRPCServer) CreateQuestion(ctx context.Context, q *quest.Question) (*quest.Result, error) {

	newQuestion := &ent.Question{ID: q.ID,
		Content:     q.Content,
		Description: q.Description,
		Answer:      q.Answer,
		UserCreated: q.UserCreated,
	}

	finalquestion, err := s.Serv.CreateQuestion(*newQuestion)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
	}

	return &quest.Result{Message: finalquestion.ID}, nil
}

func (s *QuestionGRPCServer) GetAllQuestions(context.Context, *quest.EmptyRequest) (*quest.AllQuestionResponse, error) {

	questions, err := s.Serv.GetAllQuestions()

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	questionResponse := &quest.AllQuestionResponse{}

	for _, currentQuestion := range questions {
		newQuestion := &quest.Question{}

		newQuestion.ID = currentQuestion.ID
		newQuestion.Content = currentQuestion.Content
		newQuestion.Description = currentQuestion.Description
		newQuestion.UserCreated = currentQuestion.UserCreated
		newQuestion.Answer = currentQuestion.Answer
		questionResponse.Questions = append(questionResponse.Questions, newQuestion)

	}

	return questionResponse, nil

}
func (s *QuestionGRPCServer) UpdateQuestion(ctx context.Context, req *quest.Question) (*quest.Result, error) {
	fmt.Println("Id0", req.ID)
	newQuestion := ent.Question{}
	newQuestion.ID = req.ID
	newQuestion.Content = req.Content
	newQuestion.Description = req.Description
	newQuestion.UserCreated = req.UserCreated
	newQuestion.Answer = req.Answer

	response, err := s.Serv.UpdateQuestion(req.ID, newQuestion)

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	data, _ := json.Marshal(response)

	return &quest.Result{Message: string(data)}, nil
}

func (s *QuestionGRPCServer) DeleteQuestion(ctx context.Context, req *quest.RequestWithId) (*quest.Result, error) {

	response, err := s.Serv.DeleteQuestion(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &quest.Result{Message: response}, nil
}

func (s *QuestionGRPCServer) GetQuestionById(ctx context.Context, req *quest.RequestWithId) (*quest.Question, error) {

	questionResponse, err := s.Serv.GetQuestionById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	fmt.Println("date", questionResponse.CreatedAt)

	newQuestion := &quest.Question{}

	newQuestion.ID = questionResponse.ID
	newQuestion.Content = questionResponse.Content
	newQuestion.Description = questionResponse.Description
	newQuestion.UserCreated = questionResponse.UserCreated
	newQuestion.Answer = questionResponse.Answer

	return newQuestion, nil
}
