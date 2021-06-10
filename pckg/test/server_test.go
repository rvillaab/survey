package test

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	e "survey/pckg/endpoint"
	"survey/pckg/model"
	"survey/pckg/server"
	"survey/pckg/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPServer(t *testing.T) {

	srv := mocks.NewMockService()

	getAllErr := srv.SetGetAllFunc(
		func() ([]model.Question, error) {
			questions := []model.Question{}
			questionResponse := model.Question{ID: "57", Content: "what time is it?", Answer: "it's 7 o'clock"}
			questions = append(questions, questionResponse)
			return questions, nil
		})

	if getAllErr != nil {
		log.Fatal(getAllErr)
	}

	endpoints := e.Endpoints{
		GetAllQuestionsEndpoint: e.MakeGetallQuestionsEndpoint(srv),
	}

	handler := server.NewHTTPServer(context.TODO(), endpoints)
	ts := httptest.NewServer(handler)

	defer ts.Close()

	url := ts.URL + "/questions"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bb, err1 := ioutil.ReadAll(res.Body)

	if err1 != nil {
		log.Fatal(err1)
	}

	assert.EqualValues(t, "[{\"id\":\"57\",\"content\":\"what time is it?\",\"description\":\"\",\"answer\":\"it's 7 o'clock\"}]\n", string(bb))

}
