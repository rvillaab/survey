package test

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"survey/pckg/server"
	"testing"

	"github.com/stretchr/testify/mock"
	e "survey/pckg/endpoint"
)

type QuestionServiceImplMock struct {
	mock.Mock
	DB *sql.DB
}

func TestGetAllQuestionsHandler(t testing.T) {

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}


	srv := service.NewQuestionService(enviroment.GoDotEnvVariable("DB_HOST"),
	enviroment.GoDotEnvVariable("DB_USER"),
	enviroment.GoDotEnvVariable("DB_PASSWORD"),
	enviroment.GoDotEnvVariable("DB_NAME"))

	rr := httptest.NewRecorder()
	//handler := http.HandlerFunc()
	ctx := context.Context(context.TODO())

	handler := server.NewHTTPServer(ctx, e.MakeGetallQuestionsEndpoint())
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

}
