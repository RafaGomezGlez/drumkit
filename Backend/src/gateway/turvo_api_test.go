package gateway

import (
	"testing"
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

func TestRetrieveLoads_Success2(t *testing.T) {
	repo := NewTurvoAPIGateway()

	loads, err := repo.RetrieveLoads("21", "10") // Adjust page and size as needed
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// You may want to adjust these checks based on your actual API response
	if len(loads) == 0 {
		t.Errorf("expected at least 1 load, got %d", len(loads))
	}
	// Optionally print the loads for manual inspection
	t.Logf("Retrieved loads: %+v", loads)

}
