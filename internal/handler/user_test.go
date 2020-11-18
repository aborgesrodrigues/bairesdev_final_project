package handler_test

import (
	"bairesdev_final_project/internal/domain"
	"bairesdev_final_project/internal/handler"
	"bairesdev_final_project/internal/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	birthday, _ := time.Parse("2006-01-02", "1983-01-29")
	// Data
	payload := domain.User{
		ID:       1,
		Username: "user1",
		Name:     "user name",
		Birthday: birthday,
	}
	data, marshalError := json.Marshal(payload)

	handler := handler.NewUserHandler()

	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockUserServiceInterface(ctrl)
	handler.Service.(*service.MockUserServiceInterface).EXPECT().Create(payload).Return(
		&domain.User{
			ID:       1,
			Username: "user1",
			Name:     "user name",
			Birthday: birthday,
		}, nil)

	// do the request
	req := httptest.NewRequest("POST", "/question", strings.NewReader(string(data)))
	w := httptest.NewRecorder()
	handler.CreateUser(w, req)

	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	user := domain.User{}
	unMarshalErr := json.Unmarshal(resBody, &user)

	//included to facilitate the check with the returned value
	payload.ID = 1

	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusCreated, "they should be equal")
	assert.Equal(t, user, payload, "they should be equal")
}

func TestGetAllUsers(t *testing.T) {
	allUsers := make([]domain.User, 0)
	allUsers = append(allUsers, domain.User{
		ID:       1,
		Username: "username1",
		Name:     "User Name 1",
	})
	allUsers = append(allUsers, domain.User{
		ID:       2,
		Username: "username2",
		Name:     "User Name 2",
	})

	handler := handler.NewUserHandler()
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockUserServiceInterface(ctrl)
	handler.Service.(*service.MockUserServiceInterface).EXPECT().GetAll().Return(&allUsers, nil)

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", handler.GetAllUsers).Methods("GET")

	// create request
	req := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	users := make([]domain.User, 0)
	unMarshalErr := json.Unmarshal(resBody, &users)

	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, allUsers, users, "they should be equal")
}

func TestFindByID(t *testing.T) {

	foundUser := domain.User{
		ID:       1,
		Username: "username1",
		Name:     "User Name 1",
	}

	handler := handler.NewUserHandler()
	// create the question mock interface
	ctrl := gomock.NewController(t)
	handler.Service = service.NewMockUserServiceInterface(ctrl)
	handler.Service.(*service.MockUserServiceInterface).EXPECT().FindByID(1).Return(&foundUser, nil)

	// create router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/{id}", handler.GetOneUser).Methods("GET")

	// create request
	req := httptest.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	user := domain.User{}
	unMarshalErr := json.Unmarshal(resBody, &user)

	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
	assert.Equal(t, foundUser, user, "they should be equal")
}
