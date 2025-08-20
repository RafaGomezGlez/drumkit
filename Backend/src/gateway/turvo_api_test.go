package gateway

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// mockClient implements http.Client's Do method for testing
// type mockClient struct {
// 	resp *http.Response
// 	err  error
// }

// func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
// 	return m.resp, m.err
// }

// func TestRetrieveLoads_Success(t *testing.T) {
// 	// mockLoads := `[{"id":"123","status":"active"},{"id":"456","status":"inactive"}]`
// 	// mockResp := &http.Response{
// 	// 	StatusCode: http.StatusOK,
// 	// 	Body:       ioutil.NopCloser(bytes.NewBufferString(mockLoads)),
// 	// }
// 	client := &http.Client{}
// 	repo := &TurvoAPIRepository{
// 		Host:   "https://my-sandbox-publicapi.turvo.com/v1",
// 		Client: client,
// 	}

// 	loads, err := repo.RetrieveLoads()
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}
// 	if len(loads) != 2 {
// 		t.Errorf("expected 2 loads, got %d", len(loads))
// 	}
// 	if loads[0].ID != "123" || loads[1].ID != "456" {
// 		t.Errorf("unexpected load IDs: %+v", loads)
// 	}
// }

type TurvoAPITestSuite struct {
	suite.Suite
	gw *TurvoAPIGateway
}

func (suite *TurvoAPITestSuite) SetupSuite() {
	os.Setenv("TURVO_BASE_URL", "https://my-sandbox-publicapi.turvo.com/v1")
	os.Setenv("TURVO_API_KEY", "9VjKgnIlQS1255cn7cRvJ6jNf8Z4MElP1PGgBTsH")
	os.Setenv("TURVO_CLIENT_ID", "publicapi")
	os.Setenv("TURVO_CLIENT_SECRET", "secret")
	os.Setenv("TURVO_USERNAME", "axle@wickerparklogistics.com")
	os.Setenv("TURVO_PASSWORD", "DZJ@pcu_qzd8ecz0fgw")

	suite.gw = NewTurvoAPIGateway()
	if suite.gw == nil {
		// Failed to create gateway instance
		fmt.Println("ERROR CREATING GATEWAY")
	}
}

func (suite *TurvoAPITestSuite) TestRetrieveLoads_Success() {
	loads, err := suite.gw.RetrieveLoads("21", "10")
	suite.Require().NoError(err)
	suite.NotEmpty(loads)
	suite.T().Logf("Retrieved loads: %+v", loads)
}

func TestTurvoAPITestSuite(t *testing.T) {
	suite.Run(t, new(TurvoAPITestSuite))
}
