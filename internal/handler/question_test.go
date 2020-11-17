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

	"bairesdev_final_project/internal/domain"
	"bairesdev_final_project/internal/service"

	"bairesdev_final_project/internal/handler"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestCreateQuestion tests the CreateQuestion function of the handler
func TestCreateQuestion(t *testing.T) {
	// Data
	payload := domain.Question{
		Statement: "Statement 1",
		UserID:    1,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()

	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Create(payload).Return(
		&domain.Question{
			Model:     gorm.Model{ID: 1},
			UserID:    1,
			Statement: "Statement 1",
		}, nil)

	// do the request
	req := httptest.NewRequest("POST", "/question", strings.NewReader(string(data)))
	w := httptest.NewRecorder()
	handler.CreateQuestion(w, req)

	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	question := domain.Question{}
	unMarshalErr := json.Unmarshal(resBody, &question)

	//included to facilitate the check with the returned value
	payload.ID = 1

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusCreated, "they should be equal")
	assert.Equal(t, question, payload, "they should be equal")
}

func TestFailCreateQuestion(t *testing.T) {
	// Data
	payload := domain.Question{
		Statement: "Statement 1",
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()

	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Create(payload).Return(
		&domain.Question{}, errors.New("Not all attributes are filled"))

	// do the request
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
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Update(payload).Return(
		&domain.Question{
			Model:     gorm.Model{ID: 1},
			Statement: "Statement 1",
			UserID:    1,
		}, nil)

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
	assert.Equal(t, question, payload, "they should be equal")
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
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Update(payload).Return(
		nil, nil)

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/{id}", handler.UpdateQuestion).Methods("PUT")

	// create request
	req := httptest.NewRequest("PUT", "/question/2", strings.NewReader(string(data)))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	assert.Equal(t, string(resBody), "ID of the question is different of the parameter", "they should be equal")
}

func TestFailUpdateQuestionIncompleteBody(t *testing.T) {
	// Data
	payload := domain.Question{
		Model:  gorm.Model{ID: 10},
		UserID: 1,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewQuestionHandler()
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Update(payload).Return(
		&domain.Question{}, errors.New("Not all attributes are filled"))

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
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Update(payload).Return(
		&payload, errors.New("ID not found"))

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

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	assert.Equal(t, string(resBody), "ID of the question is different of the parameter", "they should be equal")
}

func TestDeleteQuestion(t *testing.T) {
	handler := handler.NewQuestionHandler()
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Delete(1).Return(nil)

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
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().Delete(10).Return(errors.New("ID not found"))

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

func TestGetAllQuestions(t *testing.T) {
	allQuestions := make([]domain.Question, 0)
	allQuestions = append(allQuestions, domain.Question{
		Model:     gorm.Model{ID: 1},
		Statement: "statement 1",
		UserID:    1,
		User: domain.User{
			Model:    gorm.Model{ID: 1},
			Username: "username1",
			Name:     "User Name 1",
		},
	})
	allQuestions = append(allQuestions, domain.Question{
		Model:     gorm.Model{ID: 2},
		Statement: "statement 2",
		UserID:    2,
		User: domain.User{
			Model:    gorm.Model{ID: 2},
			Username: "username2",
			Name:     "User Name 2",
		},
	})

	handler := handler.NewQuestionHandler()
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().GetAll().Return(&allQuestions, nil)

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question", handler.GetAllQuestions).Methods("GET")

	// create request
	req := httptest.NewRequest("GET", "/question", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	questions := make([]domain.Question, 0)
	unMarshalErr := json.Unmarshal(resBody, &questions)

	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, allQuestions, questions, "they should be equal")
}

func TestGetQuestionsByUser(t *testing.T) {
	allUserQuestions := make([]domain.Question, 0)
	allUserQuestions = append(allUserQuestions, domain.Question{
		Model:     gorm.Model{ID: 1},
		Statement: "statement 1",
		UserID:    1,
		User: domain.User{
			Model:    gorm.Model{ID: 1},
			Username: "username1",
			Name:     "User Name 1",
		},
	})
	allUserQuestions = append(allUserQuestions, domain.Question{
		Model:     gorm.Model{ID: 2},
		Statement: "statement 2",
		UserID:    2,
		User: domain.User{
			Model:    gorm.Model{ID: 1},
			Username: "username1",
			Name:     "User Name 1",
		},
	})

	handler := handler.NewQuestionHandler()
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockQuestionServiceInterface(ctrl)
	handler.Service.(*service.MockQuestionServiceInterface).EXPECT().FindByUser(1).Return(&allUserQuestions, nil)

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question/user/{user_id}", handler.GetQuestionsByUser).Methods("GET")

	// create request
	req := httptest.NewRequest("GET", "/question/user/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	questions := make([]domain.Question, 0)
	unMarshalErr := json.Unmarshal(resBody, &questions)

	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, allUserQuestions, questions, "they should be equal")
}
