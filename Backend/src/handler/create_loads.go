package handler

import (
	"encoding/json"

	"drumkit.com/interview/src/model"
	"drumkit.com/interview/src/service"
	"github.com/aws/aws-lambda-go/events"
)

type CreateLoadsHandler struct {
	Service *service.LoadService
}

func (h *CreateLoadsHandler) CreateLoadsHandlerLambda(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Here you would typically parse the request body to get the load data
	// For simplicity, we assume the load data is valid and directly call the service
	var loads model.CreateLoadRequest
	err := json.Unmarshal([]byte(request.Body), &loads)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, err
	}

	err = h.Service.CreateLoad(loads)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Loads created successfully",
	}, nil
}
