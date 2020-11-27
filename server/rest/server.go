package server

import (
	"bairesdev_final_project/internal/service"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// ServerHandler struct
type ServerHandler struct {
	Service service.QuestionServiceInterface
	logger  *zap.Logger
}

// NewServer constructor
func NewServer() *ServerHandler {
	logger, _ := zap.NewProduction()

	service := service.NewQuestionService(logger)

	return &ServerHandler{Service: service, logger: logger}
}

// CreateRouters function
func (sh *ServerHandler) CreateRouters(router *mux.Router) {
	// mapping endpoints
	endpoints := Endpoints{
		CreateQuestion:     MakeCreateQuestionEndpoint(sh.Service),
		UpdateQuestion:     MakeUpdateQuestionEndpoint(sh.Service),
		DeleteQuestion:     MakeDeleteQuestionEndpoint(sh.Service),
		GetAllQuestion:     MakeGetAllQuestionEndpoint(sh.Service),
		GetQuestionsByUser: MakeGetQuestionsByUserpoint(sh.Service),
	}

	createQuestionHandler := kithttp.NewServer(
		endpoints.CreateQuestion,
		DecodeCreateQuestionRequest,
		EncodeCreateResponse,
		kithttp.ServerErrorEncoder(EncodeError),
	)

	updateQuestionHandler := kithttp.NewServer(
		endpoints.UpdateQuestion,
		DecodeUpdateQuestionRequest,
		EncodeResponse,
		kithttp.ServerErrorEncoder(EncodeError),
	)

	deleteQuestionHandler := kithttp.NewServer(
		endpoints.DeleteQuestion,
		DecodeIDRequest,
		EncodeResponse,
		kithttp.ServerErrorEncoder(EncodeError),
	)

	getAllHandler := kithttp.NewServer(
		endpoints.GetAllQuestion,
		decodeBlankRequest,
		EncodeResponse,
		kithttp.ServerErrorEncoder(EncodeError),
	)

	getQuestionsByUserHandler := kithttp.NewServer(
		endpoints.GetQuestionsByUser,
		DecodeIDRequest,
		EncodeResponse,
		kithttp.ServerErrorEncoder(EncodeError),
	)

	router.Methods("POST").Path("/question").Handler(createQuestionHandler)
	router.Methods("PUT").Path("/question/{id}").Handler(updateQuestionHandler)
	router.Methods("DELETE").Path("/question/{id}").Handler(deleteQuestionHandler)
	router.Methods("GET").Path("/question").Handler(getAllHandler)
	router.Methods("GET").Path("/question/user/{id}").Handler(getQuestionsByUserHandler)
}
