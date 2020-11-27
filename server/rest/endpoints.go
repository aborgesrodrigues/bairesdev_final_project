package server

import (
	"bairesdev_final_project/internal/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints struct
type Endpoints struct {
	CreateQuestion     endpoint.Endpoint
	UpdateQuestion     endpoint.Endpoint
	DeleteQuestion     endpoint.Endpoint
	GetAllQuestion     endpoint.Endpoint
	GetQuestionsByUser endpoint.Endpoint
}

// MakeCreateQuestionEndpoint function
func MakeCreateQuestionEndpoint(srv service.QuestionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCreationRequest)
		question, err := srv.Create(ctx, req.Question)

		return question, err
	}
}

// MakeUpdateQuestionEndpoint function
func MakeUpdateQuestionEndpoint(srv service.QuestionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUpdateRequest)
		question, err := srv.Update(ctx, req.Question)

		return question, err
	}
}

// MakeDeleteQuestionEndpoint function
func MakeDeleteQuestionEndpoint(srv service.QuestionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ID := request.(int)
		err := srv.Delete(ctx, ID)

		return "File deleted", err
	}
}

// MakeGetAllQuestionEndpoint function
func MakeGetAllQuestionEndpoint(srv service.QuestionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		questions, err := srv.GetAll(ctx)

		return questions, err
	}
}

// MakeGetQuestionsByUserpoint function
func MakeGetQuestionsByUserpoint(srv service.QuestionServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ID := request.(int)
		questions, err := srv.FindByUser(ctx, ID)

		return questions, err
	}
}
