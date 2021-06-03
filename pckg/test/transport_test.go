package test

import (
	"context"
	"net/http/httptest"
	"strings"
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
