package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	e "survey/pckg/endpoint"
	pb "survey/pckg/question"
	"survey/pckg/server"
	"survey/pckg/service"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		port     = flag.Int("port", 10000, "The server port")
	)

	flag.Parse()
	ctx := context.Background()
	srv := service.NewQuestionService()
	srv1 := service.NewAnswerService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := e.Endpoints{
		CountEndpoint:              e.MakeCountEndpoint(srv),
		GetAllQuestionsEndpoint:    e.MakeGetallQuestionsEndpoint(srv),
		CreateQuestionEndpoint:     e.MakeCreateQuestionEndpoint(srv),
		UpdateQuestionEndpoint:     e.MakeUpdateQuestionEndpoint(srv),
		DeleteQuestionEndpoint:     e.MakeDeleteQuestionEndpoint(srv),
		GetQuestionByIdEndpoint:    e.MakeGetQuestionByIdEndpoint(srv),
		GetQuestionsByUserEndpoint: e.MakeGetQuestionsByUserEndpoint(srv),
		GetAllAnswersEndpoint:      e.MakeGetallAnswersEndpoint(srv1),
	}

	// HTTP transport
	go func() {
		log.Println("service is listening on port:", *httpAddr)
		handler := server.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		} else {
			log.Println("service is listening on port:", *port)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterQuestionServiceServer(grpcServer, newServer(srv))
		grpcServer.Serve(lis)
	}()

	log.Fatalln(<-errChan)
}

func newServer(srv service.QuestionService) pb.QuestionServiceServer {
	s := &server.QuestionGRPCServer{Serv: srv}
	return s
}
