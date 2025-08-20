package view_loads_test

import (
	"testing"

	"drumkit.com/interview/src/gateway"
	"drumkit.com/interview/src/handler"
	"drumkit.com/interview/src/service"
	logger "drumkit.com/interview/src/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ViewLoadsTestSuite struct {
	suite.Suite
	handler *handler.ViewLoadsHandler
}

func (suite *ViewLoadsTestSuite) SetupSuite() {
	logger.NewLogger()
	gw := gateway.NewTurvoAPIGateway()
	svc := service.NewLoadService(gw)
	suite.handler = &handler.ViewLoadsHandler{Service: svc}
	// suite.T().Setenv("TOKEN", "test-token")
}

func (suite *ViewLoadsTestSuite) TestCreateLoadsHandlerLambda() {
	req := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"start":    "10",
			"pageSize": "20",
		},
	}
	resp, err := suite.handler.ViewLoadsHandlerLambda(req)
	if err != nil {
		logger.Logger.Error("Error in CreateLoadsHandlerLambda:", zap.Error(err))
	}
	suite.Equal(200, resp.StatusCode, "expected status code 200")
	logger.Logger.Info("CreateLoadsHandlerLambda executed successfully", zap.String("response", resp.Body))
}

func TestCreateLoadsTestSuite(t *testing.T) {
	suite.Run(t, new(ViewLoadsTestSuite))
}
