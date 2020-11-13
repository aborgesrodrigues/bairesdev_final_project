package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../internal/domain"
	"../internal/handler"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// TestCreateQuestion tests the CreateQuestion function of the handler
func TestCreateQuestion(t *testing.T) {
	// Data
	payload := domain.Question{
		Statement: "Statement 1",
		UserID:    1,
	}
	data, marshalError := json.Marshal(payload)

	// Create the router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/question", handler.CreateQuestion).Methods("POST")

	// Create the server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Do the call to the service
	res, err := http.Post(ts.URL+"/question", "application/json", strings.NewReader(string(data)))
	resBody, bodyErr := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	question := domain.Question{}
	unMarshalErr := json.Unmarshal(resBody, &question)

	assert.NoError(t, err)
	assert.NoError(t, marshalError)
	assert.NoError(t, bodyErr)
	assert.NoError(t, unMarshalErr)
	assert.Equal(t, res.StatusCode, http.StatusCreated, "they should be equal")
	assert.Equal(t, question.Statement, payload.Statement, "they should be equal")
	assert.Equal(t, question.UserID, payload.UserID, "they should be equal")

}
