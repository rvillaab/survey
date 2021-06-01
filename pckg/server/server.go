package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	e "survey/pckg/endpoint"
	ent "survey/pckg/model"
	quest "survey/pckg/question"
	service "survey/pckg/service"
	t "survey/pckg/transport"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		Description: q.Content,
		CreatedAt:   q.CreatedAt.AsTime().String(),
		UserCreated: q.UserCreated,
		UpdatedAt:   q.UpdatedAt.AsTime().Local().String(),
		UserUpdated: q.UserUpdated,
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

	layout := "02/01/2004 15:04:00"

	for _, currentQuestion := range questions {
		newQuestion := &quest.Question{}

		newQuestion.ID = currentQuestion.ID
		newQuestion.Content = currentQuestion.Content
		newQuestion.Description = currentQuestion.Description
		fmt.Println(currentQuestion.CreatedAt)
		crateTime, err := time.Parse(layout, currentQuestion.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		newQuestion.CreatedAt = timestamppb.New(crateTime)
		newQuestion.UserCreated = currentQuestion.UserCreated

		updateTime, erru := time.Parse(layout, currentQuestion.UpdatedAt)
		if erru != nil {
			log.Fatal(erru)
		}
		newQuestion.UpdatedAt = timestamppb.New(updateTime)
		newQuestion.UserUpdated = currentQuestion.UserUpdated

		questionResponse.Questions = append(questionResponse.Questions, newQuestion)

	}

	return questionResponse, nil

}
func (s *QuestionGRPCServer) UpdateQuestion(context.Context, *quest.RequestWithId) (*quest.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}

func (s *QuestionGRPCServer) DeleteQuestion(ctx context.Context, req *quest.RequestWithId) (*quest.Result, error) {

	response, err := s.Serv.DeleteQuestion(req.Name)

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &quest.Result{Message: response}, nil
}

func (s *QuestionGRPCServer) GetQuestionById(ctx context.Context, req *quest.RequestWithId) (*quest.Question, error) {

	questionResponse, err := s.Serv.GetQuestionById(req.Name)

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	layout := "02/01/2004 15:04:00"

	fmt.Println("date", questionResponse.CreatedAt)

	newQuestion := &quest.Question{}

	newQuestion.ID = questionResponse.ID
	newQuestion.Content = questionResponse.Content
	newQuestion.Description = questionResponse.Description
	crateTime, err := time.Parse(layout, questionResponse.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	newQuestion.CreatedAt = timestamppb.New(crateTime)
	newQuestion.UserCreated = questionResponse.UserCreated

	updateTime, erru := time.Parse(layout, questionResponse.UpdatedAt)
	if erru != nil {
		log.Fatal(erru)
	}
	newQuestion.UpdatedAt = timestamppb.New(updateTime)
	newQuestion.UserUpdated = questionResponse.UserUpdated

	return newQuestion, nil
}
