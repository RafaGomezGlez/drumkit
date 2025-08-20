package handler

import (
	"encoding/json"

	"drumkit.com/interview/src/service"
	"github.com/aws/aws-lambda-go/events"
)

type ViewLoadsHandler struct {
	Service *service.LoadService
}

func (h *ViewLoadsHandler) ViewLoadsHandlerLambda(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Call the service to retrieve loads
	// Extract query parameters from the request
	params := request.QueryStringParameters

	// Pass the query params to the service method
	loads, err := h.Service.RetrieveLoads(params["start"], params["pageSize"])
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Marshal the loads into a JSON string to be returned in the response body
	responseBody, err := json.Marshal(loads)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}
