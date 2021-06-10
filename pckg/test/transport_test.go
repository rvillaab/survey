package test

import (
	"context"
	"net/http/httptest"
	"strings"
	"survey/pckg/model"
	"survey/pckg/transport"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDecodeUpdateRequest(t *testing.T) {
	var finalResponse transport.QuestionUpdateRequest
	bodyReader := strings.NewReader(`{"ID":"2", "Content":"Who was Steve Jobs?", "Description":"About Technology", "Answer":"The man who created Apple Company"}`)
	req := httptest.NewRequest("PUT", "/questions", bodyReader)

	vars := map[string]string{
		"id": "2",
	}
	req = mux.SetURLVars(req, vars)

	response, err := transport.DecodeUpdateRequest(context.TODO(), req)

	if err != nil {
		t.Fatal(err)
	}

	finalResponse = response.(transport.QuestionUpdateRequest)

	assert.EqualValues(t, "2", finalResponse.V)
	assert.EqualValues(t, "Who was Steve Jobs?", finalResponse.S.Content)
	assert.EqualValues(t, "About Technology", finalResponse.S.Description)
	assert.EqualValues(t, "The man who created Apple Company", finalResponse.S.Answer)

}

func TestDecodeUpdateRequestErr(t *testing.T) {

	bodyReader := strings.NewReader(`{"ID":"2", "Content":"Who was Steve Jobs?", "Description":"About Technology", "Answer":"The man who created Apple Company"}`)
	req := httptest.NewRequest("PUT", "/questions", bodyReader)
	_, err := transport.DecodeUpdateRequest(context.TODO(), req)

	assert.Error(t, err)
}

func TestDecodeDeleteRequest(t *testing.T) {

	req := httptest.NewRequest("DELETE", "/questions", nil)
	_, err := transport.DecodeDeleteRequest(context.TODO(), req)
	assert.Error(t, err)

}

func TestDecodeQuestionCreateRequest(t *testing.T) {

	var finalResponse model.Question
	bodyReader := strings.NewReader(`{"ID":"25", "Content":"Firts Alphabet letter", "Description":"General", "Answer":"A"}`)
	req := httptest.NewRequest("PUT", "/questions", bodyReader)

	vars := map[string]string{
		"id": "2",
	}
	req = mux.SetURLVars(req, vars)

	response, err := transport.DecodeQuestionCreateRequest(context.TODO(), req)

	if err != nil {
		t.Fatal(err)
	}

	finalResponse = response.(model.Question)

	assert.EqualValues(t, "25", finalResponse.ID)
	assert.EqualValues(t, "Firts Alphabet letter", finalResponse.Content)
	assert.EqualValues(t, "General", finalResponse.Description)
	assert.EqualValues(t, "A", finalResponse.Answer)

}

func TestEncodeResponse(t *testing.T) {
	resp := httptest.NewRecorder()
	question := model.Question{ID: "2", Content: "test", Description: "Test", Answer: "Test"}
	transport.EncodeResponse(context.TODO(), resp, question)

	response := string(resp.Body.Bytes())
	expecResponse := "{\"id\":\"2\",\"content\":\"test\",\"description\":\"Test\",\"answer\":\"Test\"}\n"
	assert.EqualValues(t, expecResponse, response)

}
