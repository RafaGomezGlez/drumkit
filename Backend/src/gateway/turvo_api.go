package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"drumkit.com/interview/src/model"
)

type AuthRequestBody struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Scope     string `json:"scope"`
	Type      string `json:"type"`
}

type AuthResponseBody struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type TurvoAPIGateway struct {
	Host         string
	Client       httpClient
	Token        string
	APIKey       string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

// getAuthToken is a method so it can use values stored on the gateway instance.
func (r *TurvoAPIGateway) getAuthToken() (string, error) {
	tokenURL := r.Host + "/oauth/token"
	requestBody, err := json.Marshal(AuthRequestBody{
		GrantType: "password",
		Username:  r.Username,
		Password:  r.Password,
		Scope:     "read+trust+write",
		Type:      "business",
	})
	if err != nil {
		return "", fmt.Errorf("error creating request body JSON: %w", err)
	}

	req, err := http.NewRequest("POST", tokenURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating HTTP request: %w", err)
	}

	req.Header.Set("x-api-key", r.APIKey)
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("client_id", r.ClientID)
	q.Add("client_secret", r.ClientSecret)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending token request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading token response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var tokenResponse AuthResponseBody
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", fmt.Errorf("error parsing token response JSON: %w", err)
	}

	if tokenResponse.AccessToken == "" {
		return "", fmt.Errorf("access token was empty in the response")
	}

	return tokenResponse.AccessToken, nil
}

// NewTurvoAPIGateway reads credentials from environment variables (with sensible defaults)
// and initializes the gateway instance. This avoids package-level globals.
func NewTurvoAPIGateway() *TurvoAPIGateway {
	host := os.Getenv("TURVO_BASE_URL")
	if host == "" {
		return nil
	}
	apiKey := os.Getenv("TURVO_API_KEY")
	if apiKey == "" {
		return nil
	}
	clientID := os.Getenv("TURVO_CLIENT_ID")
	if clientID == "" {
		return nil
	}
	clientSecret := os.Getenv("TURVO_CLIENT_SECRET")
	if clientSecret == "" {
		return nil
	}
	username := os.Getenv("TURVO_USERNAME")
	if username == "" {
		return nil
	}
	password := os.Getenv("TURVO_PASSWORD")
	if password == "" {
		return nil
	}

	gw := &TurvoAPIGateway{
		Host:         host,
		Client:       &http.Client{},
		APIKey:       apiKey,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
	}

	token, err := gw.getAuthToken()
	if err != nil {
		panic(fmt.Sprintf("failed to authenticate: %v", err))
	}
	gw.Token = token

	return gw
}
func (r *TurvoAPIGateway) RetrieveLoads(start string, pageSize string) ([]model.Shipment, error) {
	u, err := url.Parse(fmt.Sprintf("%s/shipments/list", r.Host))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if start != "" {
		q.Set("start", start)
	}
	if pageSize != "" {
		q.Set("pageSize", pageSize)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+r.Token)
	req.Header.Set("x-api-key", r.APIKey)

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var shipmentsResp model.ShipmentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&shipmentsResp); err != nil {
		return nil, err
	}

	return shipmentsResp.Details.Shipments, nil
}

func (r *TurvoAPIGateway) CreateLoad(loads []model.CreateLoadRequest) error {
	for _, load := range loads {
		avroLoadRequest, err := transformCreateLoadRequestToAvro(load)
		if err != nil {
			return fmt.Errorf("failed to transform load request: %w", err)
		}

		requestBody, err := json.Marshal(avroLoadRequest)
		if err != nil {
			return fmt.Errorf("error marshalling AvroLoadRequest: %w", err)
		}

		url := fmt.Sprintf("%s/shipments", r.Host)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			return fmt.Errorf("error creating HTTP request: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+r.Token)
		req.Header.Set("x-api-key", r.APIKey)
		req.Header.Set("Content-Type", "application/json")

		resp, err := r.Client.Do(req)
		if err != nil {
			return fmt.Errorf("error sending create load request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("failed to create load, status code: %d, response: %s", resp.StatusCode, string(body))
		}
	}
	return nil
}

// transformCreateLoadRequestToAvro transforms the simplified CreateLoadRequest
// into the detailed AvroLoadRequest required by the Turvo API.
func transformCreateLoadRequestToAvro(input model.CreateLoadRequest) (model.AvroLoadRequest, error) {

	// Parse the dates from the input, assuming they are in a string format
	startTime, err := time.Parse(time.RFC3339, "2025-08-22T09:00:00Z")
	if err != nil {
		return model.AvroLoadRequest{}, fmt.Errorf("failed to parse start time: %w", err)
	}

	endTime, err := time.Parse(time.RFC3339, "2025-08-25T17:00:00Z")
	if err != nil {
		return model.AvroLoadRequest{}, fmt.Errorf("failed to parse end time: %w", err)
	}

	avroRequest := model.AvroLoadRequest{
		LTLShipment: false,
		StartDate: model.DateTime{
			Date:     startTime.Format(time.RFC3339),
			TimeZone: "America/Chicago",
		},
		EndDate: model.DateTime{
			Date:     endTime.Format(time.RFC3339),
			TimeZone: "America/Chicago",
		},
		Status: model.AvroStatus{
			Code: model.ValueKey{
				Key:   "2102",
				Value: "Covered",
			},
		},
		Lane: model.Lane{
			Start: "Chicago, IL",
			End:   "Dallas, TX",
		},
		GlobalRoute: []model.GlobalRoute{
			{
				StopType: model.ValueKey{Key: "1500", Value: "Pickup"},
				Location: model.Location{ID: 525817},
				Sequence: 0,
				Appointment: model.Appointment{
					Date:     startTime.Format(time.RFC3339),
					Flex:     3600,
					Timezone: "America/Chicago",
					HasTime:  true,
				},
				CustomerOrder: []model.GlobalRouteCustomerOrder{{CustomerID: 973069, CustomerOrderSourceID: 1}},
				CarrierOrder:  []model.GlobalRouteCarrierOrder{{CarrierID: 973069, CarrierOrderSourceID: 1}},
			},
			{
				StopType: model.ValueKey{Key: "1501", Value: "Delivery"},
				Location: model.Location{ID: 525817},
				Sequence: 1,
				Appointment: model.Appointment{
					Date:     endTime.Format(time.RFC3339),
					Flex:     3600,
					Timezone: "America/Chicago",
					HasTime:  true,
				},
				CustomerOrder: []model.GlobalRouteCustomerOrder{{CustomerID: 973069, CustomerOrderSourceID: 1}},
				CarrierOrder:  []model.GlobalRouteCarrierOrder{{CarrierID: 973069, CarrierOrderSourceID: 1}},
			},
		},
		CustomerOrder: []model.CustomerOrderAvro{
			{
				CustomerOrderSourceID: 1,
				Customer: model.CustomerAvro{
					ID: 973069,
				},
				Items: []model.Item{
					{
						Name:      "Proof of Concept Item",
						Qty:       10,
						Unit:      model.ValueKey{Key: "6000", Value: "pallets"},
						Operation: 0,
					},
				},
				Costs: model.Costs{
					TotalAmount: 1500.00,
					LineItem: []model.LineItem{
						{
							Code:      model.ValueKey{Key: "1600", Value: "Freight - flat"},
							Qty:       1,
							Price:     1500.00,
							Amount:    1500.00,
							Billable:  true,
							Operation: 0,
						},
					},
				},
			},
		},
	}
	return avroRequest, nil
}
