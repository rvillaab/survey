package server

import (
	"context"
	"net/http"
	e "survey/pckg/endpoint"
	ent "survey/pckg/entity"
	quest "survey/pckg/question"
	service "survey/pckg/service"
	t "survey/pckg/transport"

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

	r.Methods("GET").Path("/count").Handler(httptransport.NewServer(
		endpoints.CountEndpoint,
		t.DecodeCountRequest,
		t.EncodeResponse,
	))

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
		CreatedAt:   q.CreatedAt.AsTime(),
		UserCreated: q.UserCreated,
		UpdatedAt:   q.UpdatedAt.AsTime(),
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

	for _, currentQuestion := range questions {
		newQuestion := &quest.Question{ID: currentQuestion.ID,
			Content:     currentQuestion.Content,
			Description: currentQuestion.Description,
			CreatedAt:   timestamppb.New(currentQuestion.CreatedAt),
			UserCreated: currentQuestion.UserCreated,
			UpdatedAt:   timestamppb.New(currentQuestion.UpdatedAt),
			UserUpdated: currentQuestion.UserUpdated}

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

	newQuestion := &quest.Question{ID: questionResponse.ID,
		Content:     questionResponse.Content,
		Description: questionResponse.Content,
		CreatedAt:   timestamppb.New(questionResponse.CreatedAt),
		UserCreated: questionResponse.UserCreated,
		UpdatedAt:   timestamppb.New(questionResponse.UpdatedAt),
		UserUpdated: questionResponse.UserUpdated,
	}

	return newQuestion, nil
}
