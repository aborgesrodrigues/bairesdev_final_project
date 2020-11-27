package server

import (
	"bairesdev_final_project/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ErrorResponse struct
type ErrorResponse struct {
	Err string `json:"error,omitempty"`
}

// GetCreationRequest struct
type GetCreationRequest struct {
	domain.Question
}

// GetUpdateRequest struct
type GetUpdateRequest struct {
	domain.Question
}

func decodeBlankRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

// DecodeCreateQuestionRequest function
func DecodeCreateQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
	var request GetCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return nil, err
	}

	return request, nil
}

// DecodeUpdateQuestionRequest function
func DecodeUpdateQuestionRequest(c context.Context, r *http.Request) (interface{}, error) {
	var request GetUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return nil, err
	}

	ID := mux.Vars(r)["id"]
	intID, intError := strconv.Atoi(ID)

	if intError != nil {
	}

	if intID != int(request.ID) {
		return domain.Question{}, errors.New("client: URI and Body payload ids do not match")
	}

	return request, nil
}

// DecodeIDRequest function
func DecodeIDRequest(c context.Context, r *http.Request) (interface{}, error) {
	ID := mux.Vars(r)["id"]
	intID, intError := strconv.Atoi(ID)

	if intError != nil {
		return nil, intError
	}

	return intID, nil
}

// EncodeResponse function
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// EncodeCreateResponse function
func EncodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

// EncodeError function
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	statusCode, msg := DecodeError(err.Error())
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Err: msg})
}

// DecodeError function
func DecodeError(err string) (int, string) {
	return http.StatusBadRequest, err
}
