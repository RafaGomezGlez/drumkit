package create_loads_test

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

type CreateLoadsTestSuite struct {
	suite.Suite
	handler *handler.CreateLoadsHandler
}

func (suite *CreateLoadsTestSuite) SetupSuite() {
	logger.NewLogger()
	gw := gateway.NewTurvoAPIGateway()
	svc := service.NewLoadService(gw)
	suite.handler = &handler.CreateLoadsHandler{Service: svc}
	// suite.T().Setenv("TOKEN", "test-token")
}

func (suite *CreateLoadsTestSuite) TestCreateLoadsHandlerLambda() {
	req := events.APIGatewayProxyRequest{
		Body: `{
			"externalTMSLoadID": "string",
			"freightLoadID": "string",
			"status": "Quote",
			"customer": {
				"externalTMSId": "string",
				"name": "string",
				"addressLine1": "string",
				"addressLine2": "string",
				"city": "string",
				"state": "string",
				"zipcode": "string",
				"country": "string",
				"contact": "string",
				"phone": "string",
				"email": "string",
				"refNumber": "string"
				},
			"billTo": {
				"externalTMSId": "string",
				"name": "string",
				"addressLine1": "string",
				"addressLine2": "string",
				"city": "string",
				"state": "string",
				"zipcode": "string",
				"country": "string",
				"contact": "string",
				"phone": "string",
				"email": "string"
				},
			"pickup": {
				"externalTMSId": "string",
				"name": "string",
				"addressLine1": "string",
				"addressLine2": "string",
				"city": "string",
				"state": "string",
				"zipcode": "string",
				"country": "string",
				"contact": "string",
				"phone": "string",
				"email": "string",
				"businessHours": "string",
				"refNumber": "string",
				"readyTime": "2025-08-19T23:31:45.558Z",
				"apptTime": "2025-08-19T23:31:45.558Z",
				"apptNote": "string",
				"timezone": "string",
				"warehouseId": "string"
			},
			"consignee": {
				"externalTMSId": "string",
				"name": "string",
				"addressLine1": "string",
				"addressLine2": "string",
				"city": "string",
				"state": "string",
				"zipcode": "string",
				"country": "string",
				"contact": "string",
				"phone": "string",
				"email": "string",
				"businessHours": "string",
				"refNumber": "string",
				"mustDeliver": "string",
				"apptTime": "2025-08-19T23:31:45.558Z",
				"apptNote": "string",
				"timezone": "string",
				"warehouseId": "string"
			},
			"carrier": {
				"mcNumber": "string",
				"dotNumber": "string",
				"name": "string",
				"phone": "string",
				"dispatcher": "string",
				"sealNumber": "string",
				"scac": "string",
				"firstDriverName": "string",
				"firstDriverPhone": "string",
				"secondDriverName": "string",
				"secondDriverPhone": "string",
				"email": "string",
				"dispatchCity": "string",
				"dispatchState": "string",
				"externalTMSTruckId": "string",
				"externalTMSTrailerId": "string",
				"confirmationSentTime": "2025-08-19T23:31:45.558Z",
				"confirmationReceivedTime": "2025-08-19T23:31:45.558Z",
				"dispatchedTime": "2025-08-19T23:31:45.558Z",
				"expectedPickupTime": "2025-08-19T23:31:45.558Z",
				"pickupStart": "2025-08-19T23:31:45.558Z",
				"pickupEnd": "2025-08-19T23:31:45.558Z",
				"expectedDeliveryTime": "2025-08-19T23:31:45.558Z",
				"deliveryStart": "2025-08-19T23:31:45.558Z",
				"deliveryEnd": "2025-08-19T23:31:45.558Z",
				"signedBy": "string",
				"externalTMSId": "string"
			},
			"rateData": {
				"customerRateType": "string",
				"customerNumHours": 0,
				"customerLhRateUsd": 0,
				"fscPercent": 0,
				"fscPerMile": 0,
				"carrierRateType": "string",
				"carrierNumHours": 0,
				"carrierLhRateUsd": 0,
				"carrierMaxRate": 0,
				"netProfitUsd": 0,
				"profitPercent": 0
			},
			"specifications": {
				"minTempFahrenheit": 0,
				"maxTempFahrenheit": 0,
				"liftgatePickup": true,
				"liftgateDelivery": true,
				"insidePickup": true,
				"insideDelivery": true,
				"tarps": true,
				"oversized": true,
				"hazmat": true,
				"straps": true,
				"permits": true,
				"escorts": true,
				"seal": true,
				"customBonded": true,
				"labor": true
			},
				"inPalletCount": 0,
				"outPalletCount": 0,
				"numCommodities": 0,
				"totalWeight": 0,
				"billableWeight": 0,
				"poNums": "string",
				"operator": "string",
				"routeMiles": 0
			}`,
	}
	resp, err := suite.handler.CreateLoadsHandlerLambda(req)
	if err != nil {
		logger.Logger.Error("Error in CreateLoadsHandlerLambda:", zap.Error(err))
	}
	suite.Equal(200, resp.StatusCode, "expected status code 200")
	logger.Logger.Info("CreateLoadsHandlerLambda executed successfully", zap.String("response", resp.Body))
}

func TestCreateLoadsTestSuite(t *testing.T) {
	suite.Run(t, new(CreateLoadsTestSuite))
}
