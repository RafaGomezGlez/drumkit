package service

import (
	"os"
	"testing"

	"drumkit.com/interview/src/gateway"
	"drumkit.com/interview/src/model"
	"github.com/stretchr/testify/suite"
)

type LoadServiceTestSuite struct {
	suite.Suite
	service *LoadService
}

func (suite *LoadServiceTestSuite) SetupSuite() {
	os.Setenv("TURVO_BASE_URL", "https://my-sandbox-publicapi.turvo.com/v1")
	os.Setenv("TURVO_API_KEY", "9VjKgnIlQS1255cn7cRvJ6jNf8Z4MElP1PGgBTsH")
	os.Setenv("TURVO_CLIENT_ID", "publicapi")
	os.Setenv("TURVO_CLIENT_SECRET", "secret")
	os.Setenv("TURVO_USERNAME", "axle@wickerparklogistics.com")
	os.Setenv("TURVO_PASSWORD", "DZJ@pcu_qzd8ecz0fgw")

	gw := gateway.NewTurvoAPIGateway()
	suite.Require().NotNil(gw)
	suite.service = NewLoadService(gw)
}

func (suite *LoadServiceTestSuite) TestCreateLoad_Success() {
	loadReq := model.CreateLoadRequest{
		Pickup: model.CLPickup{
			Name:     "test",
			ApptTime: "2023-10-01T08:00:00Z",
			City:     "Chicago",
			State:    "IL",
			Country:  "USA",
		},
		Consignee: model.CLConsignee{
			Name:     "test",
			ApptTime: "2023-10-01T17:00:00Z",
			City:     "Los Angeles",
			State:    "CA",
			Country:  "USA",
		},
		Status: "Covered",
		Specifications: model.CLSpecifications{
			MinTempFahrenheit: 32,
			MaxTempFahrenheit: 75,
		},
		TotalWeight: 10000,
	}
	err := suite.service.CreateLoad(loadReq)
	suite.NoError(err)
}

func TestLoadServiceTestSuite(t *testing.T) {
	suite.Run(t, new(LoadServiceTestSuite))
}
