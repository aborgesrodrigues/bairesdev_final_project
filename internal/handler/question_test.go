package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../domain"
	"../handler"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type mockQuestionService struct{}

func (mqs *mockQuestionService) Create(question domain.Question) (*domain.Question, error) {
	if question.Statement == "" ||
		question.UserID == 0 {

		return &domain.Question{}, errors.New("Not all attributes are filled")
	}

	question.ID = 1
	return &question, nil
}

func (mqs *mockQuestionService) Update(question domain.Question) (*domain.Question, error) {
	if question.ID == 0 ||
		question.Statement == "" ||
		question.UserID == 0 {

		return nil, errors.New("Not all attributes are filled")
	}

	if question.ID == 10 {
		return nil, errors.New("ID not found")
	}

	return &question, nil
}

func (mqs *mockQuestionService) Delete(ID int) error {
	if ID == 10 {
		return errors.New("ID not found")
	}

	return nil
}

func (mqs *mockQuestionService) FindByID(ID int) (*domain.Question, error) {
	if ID == 10 {
		return nil, errors.New("ID not found")
	}

	return &domain.Question{
		Model:     gorm.Model{ID: 1},
		Statement: "Statement 1",
		UserID:    1,
	}, nil
}

func (mqs *mockQuestionService) FindByUser(userID int) (*[]domain.Question, error) {
	var questions = make([]domain.Question, 0)

	return &questions, nil
}

func (mqs *mockQuestionService) GetAll() (*[]domain.Question, error) {
	var questions = make([]domain.Question, 0)

	return &questions, nil
}

func createMockService() *mockQuestionService {
	return &mockQuestionService{}
}

// TestCreateQuestion tests the CreateQuestion function of the handler
func TestCreateQuestion(t *testing.T) {
	// Data
	payload := domain.Question{
		Statement: "Statement 1",
		UserID:    1,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	req := httptest.NewRequest("POST", "/question", strings.NewReader(string(data)))
	w := httptest.NewRecorder()
	handler.CreateQuestion(w, req)

	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	question := domain.Question{}
	unMarshalErr := json.Unmarshal(resBody, &question)

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusCreated, "they should be equal")
	assert.Equal(t, question.Statement, payload.Statement, "they should be equal")
	assert.Equal(t, question.UserID, payload.UserID, "they should be equal")
}

func TestFailCreateQuestion(t *testing.T) {
	// Data
	payload := domain.Question{
		Statement: "Statement 1",
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	req := httptest.NewRequest("POST", "/question", strings.NewReader(string(data)))
	w := httptest.NewRecorder()
	handler.CreateQuestion(w, req)

	res := w.Result()
	defer res.Body.Close()
	_, bodyErr := ioutil.ReadAll(res.Body)

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
}

func TestUpdateQuestion(t *testing.T) {
	// Data
	payload := domain.Question{
		Model:     gorm.Model{ID: 1},
		Statement: "Statement 1",
		UserID:    1,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/{id}", handler.UpdateQuestion).Methods("PUT")

	// create request
	req := httptest.NewRequest("PUT", "/question/1", strings.NewReader(string(data)))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	question := domain.Question{}
	unMarshalErr := json.Unmarshal(resBody, &question)

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, question.Statement, payload.Statement, "they should be equal")
	assert.Equal(t, question.UserID, payload.UserID, "they should be equal")
}

func TestFailUpdateQuestionDifferentID(t *testing.T) {
	// Data
	payload := domain.Question{
		Model:     gorm.Model{ID: 1},
		Statement: "Statement 1",
		UserID:    1,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/{id}", handler.UpdateQuestion).Methods("PUT")

	// create request
	req := httptest.NewRequest("PUT", "/question/2", strings.NewReader(string(data)))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	_, bodyErr := ioutil.ReadAll(res.Body)

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
}

func TestFailUpdateQuestionIncompleteBody(t *testing.T) {
	// Data
	payload := domain.Question{
		Model:  gorm.Model{ID: 10},
		UserID: 1,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/{id}", handler.UpdateQuestion).Methods("PUT")

	// create request
	req := httptest.NewRequest("PUT", "/question/10", strings.NewReader(string(data)))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	assert.Equal(t, string(resBody), "Not all attributes are filled", "they should be equal")
}

func TestFailUpdateQuestionNotFound(t *testing.T) {
	// Data
	payload := domain.Question{
		Model:     gorm.Model{ID: 10},
		Statement: "Statement 1",
		UserID:    1,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/{id}", handler.UpdateQuestion).Methods("PUT")

	// create request
	req := httptest.NewRequest("PUT", "/question/1", strings.NewReader(string(data)))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)
	fmt.Println(resBody)

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	assert.Equal(t, string(resBody), "ID of the question is different of the parameter", "they should be equal")
}

func TestDeleteQuestion(t *testing.T) {
	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/{id}", handler.DeleteQuestion).Methods("DELETE")

	// create request
	req := httptest.NewRequest("DELETE", "/question/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)
	fmt.Println(resBody)

	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
}

func TestFailDeleteQuestion(t *testing.T) {
	handler := handler.NewQuestionHandler()
	handler.Service = createMockService()

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/{id}", handler.DeleteQuestion).Methods("DELETE")

	// create request
	req := httptest.NewRequest("DELETE", "/question/10", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)
	fmt.Println(resBody)

	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
}
