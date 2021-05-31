package server

import (
	"context"
	"fmt"
	"net/http"
	e "survey/pckg/endpoint"
	quest "survey/pckg/question"
	service "survey/pckg/service"
	t "survey/pckg/transport"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
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

func (s *QuestionGRPCServer) GetQuestions(ctx context.Context, name *quest.Request) (*quest.Result, error) {
	service := s.Serv
	return &quest.Result{Message: fmt.Sprint(service.Count())}, nil
	//return &quest.Result{}, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
