package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	ent "survey/pckg/entity"

	"github.com/gorilla/mux"
)

type QuestionUpdateRequest struct {
	S ent.Question `json:"s"`
	V string       `json:"v"`
}

type QuestionRequest struct {
	Str string `json:"str"`
}

type QuestionResponse struct {
	Str string `json:"str"`
	Err string `json:"err,omitempty"`
}

type QuestionsResponse struct {
	Str string
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeQuestionCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ent.Question
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func DecodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request QuestionUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request.S); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}

	request.V = id
	return request, nil
}

func DecodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request QuestionRequest

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("id is missing in parameters")
	}

	request.Str = id
	return request, nil
}

func DecodeGetByUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request QuestionRequest

	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		return nil, errors.New("name is missing in parameters")
	}

	request.Str = name
	return request, nil
}
