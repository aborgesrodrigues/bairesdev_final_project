package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"bairesdev_final_project/internal/domain"
	"bairesdev_final_project/internal/service"
)

// UserHandler struct
type UserHandler struct {
	Service service.UserServiceInterface
	logger  *zap.Logger
}

// NewUserHandler constructor to QuestionHandler struct
func NewUserHandler() *UserHandler {
	logger, _ := zap.NewProduction()
	return &UserHandler{Service: service.NewUserService(logger), logger: logger}
}

//CreateUser func
func (qu *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// log
	defer qu.logger.Sync()

	var user domain.User
	reqBody, err := ioutil.ReadAll(r.Body)

	// Entry log
	qu.logger.Info("Fetched CreateUser",
		zap.String("url", r.URL.Path),
		zap.String("body", string(reqBody)),
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data with the Product title and description only in order to update")

		//log
		qu.logger.Error(err.Error(),
			zap.String("url", r.URL.Path),
			zap.String("body", string(reqBody)),
		)
		return
	}
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)

		//log
		qu.logger.Error(err.Error(),
			zap.String("url", r.URL.Path),
			zap.String("body", string(reqBody)),
		)
		return
	}

	createdUser, error := qu.Service.Create(user)

	if error != nil {
		fmt.Fprintf(w, error.Error())

		//log
		qu.logger.Error(error.Error(),
			zap.String("url", r.URL.Path),
			zap.String("body", string(reqBody)),
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// GetOneUser func
func (qu *UserHandler) GetOneUser(w http.ResponseWriter, r *http.Request) {
	// log
	defer qu.logger.Sync()

	// Entry log
	qu.logger.Info("Fetched GetQuestionsByUser",
		zap.String("url", r.URL.Path),
		zap.String("variables", fmt.Sprintf("%#v", mux.Vars(r))),
	)

	userID := mux.Vars(r)["id"]
	intUserID, intError := strconv.Atoi(userID)

	if intError != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid number")

		//log
		qu.logger.Error(intError.Error(),
			zap.String("url", r.URL.Path),
			zap.String("variables", fmt.Sprintf("%#v", mux.Vars(r))),
		)
		return
	}

	user, error := qu.Service.FindByID(intUserID)

	if error != nil {
		fmt.Fprintf(w, error.Error())

		//log
		qu.logger.Error(error.Error(),
			zap.String("url", r.URL.Path),
			zap.String("variables", fmt.Sprintf("%#v", mux.Vars(r))),
		)
		return
	}

	json.NewEncoder(w).Encode(user)

}

// GetAllUsers func
func (qu *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// log
	defer qu.logger.Sync()

	// Entry log
	qu.logger.Info("Fetched GetQuestionsByUser",
		zap.String("url", r.URL.Path),
	)

	users, error := qu.Service.GetAll()

	if error != nil {
		fmt.Fprintf(w, error.Error())

		//log
		qu.logger.Error(error.Error(),
			zap.String("url", r.URL.Path),
			zap.String("variables", fmt.Sprintf("%#v", mux.Vars(r))),
		)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// CreateUserRouters func
func CreateUserRouters(router *mux.Router) {
	qu := NewUserHandler()

	router.HandleFunc("/user", qu.CreateUser).Methods("POST")
	router.HandleFunc("/user", qu.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", qu.GetOneUser).Methods("GET")
}
