package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"drumkit.com/interview/src/model"
	"drumkit.com/interview/src/model/customer"
	locations "drumkit.com/interview/src/model/location"
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
func (r *TurvoAPIGateway) RetrieveLocations(query string) ([]locations.Location, error) {
	u, err := url.Parse(fmt.Sprintf("%s/locations/list", r.Host))
	if err != nil {
		return nil, err
	}

	if query != "" {
		// keep the literal brackets in the final URL, escape the inner value
		u.RawQuery = fmt.Sprintf("name[in]=[%s]", url.QueryEscape(query))
	} else {
		u.RawQuery = ""
	}

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

	var locationsResp locations.LocationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationsResp); err != nil {
		return nil, err
	}

	return locationsResp.Details.Locations, nil
}

func (r *TurvoAPIGateway) RetrieveCustomers(query string) ([]customer.Customer, error) {
	u, err := url.Parse(fmt.Sprintf("%s/customers/list", r.Host))
	if err != nil {
		return nil, err
	}

	if query != "" {
		u.RawQuery = fmt.Sprintf("name[in]=[%s]", url.QueryEscape(query))
	} else {
		u.RawQuery = ""
	}

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

	var customersResp customer.CustomersResponse
	if err := json.NewDecoder(resp.Body).Decode(&customersResp); err != nil {
		return nil, err
	}

	return customersResp.Details.Customers, nil
}

func (r *TurvoAPIGateway) CreateLoad(loads []model.CreateLoadRequest, pickUpId, deliveryId int) error {
	for _, load := range loads {
		avroLoadRequest, err := transformCreateLoadRequestToAvro(load, pickUpId, deliveryId)
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

func transformCreateLoadRequestToAvro(input model.CreateLoadRequest, pickUpId, deliveryId int) (model.AvroLoadRequest, error) {

	pickupTime, err := time.Parse(time.RFC3339, input.Pickup.ApptTime)
	if err != nil {
		return model.AvroLoadRequest{}, fmt.Errorf("failed to parse pickup apptTime: %w", err)
	}
	consigneeTime, err := time.Parse(time.RFC3339, input.Consignee.ApptTime)
	if err != nil {
		return model.AvroLoadRequest{}, fmt.Errorf("failed to parse consignee apptTime: %w", err)
	}
	statusCode, ok := model.StatusCodeForValue(input.Status)
	if !ok {
		return model.AvroLoadRequest{}, fmt.Errorf("invalid status code: %s", input.Status)

	}
	ltlShipment := false
	if input.TotalWeight < 15000 {
		ltlShipment = true
	}

	// determine equipment type from specifications
	equipTypeKey := "1200"
	equipTypeValue := "Van"
	isReefer := false

	if input.Specifications.MinTempFahrenheit != 0 || input.Specifications.MaxTempFahrenheit != 0 {
		equipTypeKey = "1208"
		equipTypeValue = "Refrigerated"
		isReefer = true
	} else if input.Specifications.Tarps || input.Specifications.Oversized {
		equipTypeKey = "1204"
		equipTypeValue = "Flatbed"
	}

	// build equipment entry
	equipment := model.Equipment{
		Operation: 0,
		Type:      model.ValueKey{Key: equipTypeKey, Value: equipTypeValue},
		Size:      model.ValueKey{Key: "1000", Value: "53ft"},
		Weight:    int(input.TotalWeight),
		WeightUnits: model.ValueKey{
			Key:   "1520",
			Value: "lb",
		},
	}

	// if reefer, include temp fields
	if isReefer {
		// prefer minTemp if present, otherwise maxTemp
		var tempVal float64
		if input.Specifications.MinTempFahrenheit != 0 {
			tempVal = float64(input.Specifications.MinTempFahrenheit)
		} else {
			tempVal = float64(input.Specifications.MaxTempFahrenheit)
		}
		// set temp and temp units if the model supports them
		// (Temp and TempUnits fields must exist on model.Equipment)
		equipment.Temp = int(tempVal)
		equipment.TempUnits = model.ValueKey{Key: "1510", Value: "°F"}
	}
	laneStart := input.Pickup.City + ", " + input.Pickup.State
	laneEnd := input.Consignee.City + ", " + input.Consignee.State

	// transform poNums -> []string
	poNumbers := splitPONumbers(input.PoNums)

	// build global route pickup stop
	pickupStop := model.GlobalRoute{
		StopType: model.ValueKey{Key: "1500", Value: "Pickup"},
		Name:     input.Pickup.Name,
		Location: model.Location{ID: pickUpId},
		Sequence: 0,
		Appointment: model.Appointment{
			Date:     pickupTime.Format(time.RFC3339),
			Flex:     3600,
			Timezone: input.Pickup.Timezone,
			HasTime:  true,
		},
		PoNumbers:     poNumbers,
		Notes:         input.Pickup.ApptNote,
		CustomerOrder: []model.GlobalRouteCustomerOrder{}, // map if you have customerOrder references in input
		CarrierOrder:  []model.GlobalRouteCarrierOrder{},  // map if you have carrierOrder references in input
	}

	// build global route delivery stop (basic mapping)
	_ = model.GlobalRoute{
		StopType: model.ValueKey{Key: "1501", Value: "Delivery"},
		Name:     input.Consignee.Name,
		Location: model.Location{ID: deliveryId},
		Sequence: 1,
		Appointment: model.Appointment{
			Date:     consigneeTime.Format(time.RFC3339),
			Flex:     3600,
			Timezone: input.Consignee.Timezone,
			HasTime:  true,
		},
		PoNumbers:     poNumbers,
		Notes:         input.Consignee.ApptNote,
		CustomerOrder: []model.GlobalRouteCustomerOrder{},
		CarrierOrder:  []model.GlobalRouteCarrierOrder{},
	}
	// Convert ExternalTMSId from string to int
	externalTMSId := 0
	if input.Customer.ExternalTMSId != "" {
		if id, err := strconv.Atoi(input.Customer.ExternalTMSId); err == nil {
			externalTMSId = id
		}
	}
	// build externalIds from poNumbers -> []model.ExternalID
	var externalIDs []model.ExternalID
	if len(poNumbers) > 0 {
		externalIDs = make([]model.ExternalID, 0, len(poNumbers))
		for _, p := range poNumbers {
			externalIDs = append(externalIDs, model.ExternalID{
				// type indicates this is a purchase order number in Turvo
				Type:  model.ValueKey{Key: "1400", Value: "Purchase order #"},
				Value: p,
			})
		}
	}

	avroRequest := model.AvroLoadRequest{
		LTLShipment: ltlShipment,
		StartDate: model.DateTime{
			Date:     pickupTime.Format("2006-01-02T15:04:05Z"),
			TimeZone: "America/Chicago",
		},
		EndDate: model.DateTime{
			Date:     consigneeTime.Format("2006-01-02T15:04:05Z"),
			TimeZone: "America/Chicago",
		},

		Status: model.AvroStatus{
			Description: input.Status,
			Notes:       "Created via API", // TODO Check if this is needed
			Code: model.ValueKey{
				Key:   string(statusCode), // Assuming "2102" is the code for "Covered"
				Value: input.Status,
			},
		},
		Groups: []model.Group{
			{
				ID:        7839, // Using existing group!
				Name:      "Drumkit Test",
				Operation: 0, // INSERT operation
			},
		},
		// ! Contributors: []model.Contributor will not be used!
		Lane:        model.Lane{Start: laneStart, End: laneEnd},
		Equipment:   []model.Equipment{equipment},
		GlobalRoute: []model.GlobalRoute{pickupStop},
		CustomerOrder: []model.CustomerOrderAvro{
			{
				CustomerOrderSourceID: externalTMSId,
				Customer: model.CustomerAvro{
					ID:   externalTMSId,
					Name: input.Customer.Name,
				},
				Items: []model.Item{
					{
						Name:         "Proof of Concept Item",
						HandlingQty:  input.InPalletCount,
						HandlingUnit: model.ValueKey{Key: "35210", Value: "Pallets"},
						IsHazmat:     input.Specifications.Hazmat,
						MinTemp: model.Temperature{
							Temp: input.Specifications.MinTempFahrenheit,
							TempUnit: model.ValueKey{
								Key:   "1510",
								Value: "°F",
							},
						},
						MaxTemp: model.Temperature{
							Temp: input.Specifications.MaxTempFahrenheit,
							TempUnit: model.ValueKey{
								Key:   "1510",
								Value: "°F",
							},
						},
					},
				},
				Costs: model.Costs{
					TotalAmount: input.RateData.CustomerLhRateUsd,
					LineItem: []model.LineItem{
						{

							Code:      model.ValueKey{Key: "1600", Value: "Freight - flat"},
							Price:     input.RateData.CustomerLhRateUsd,
							Amount:    input.RateData.CustomerLhRateUsd,
							Billable:  true,
							Operation: 0,
						},
					},
				},
				ExternalIds: externalIDs,
			},
		},
	}
	return avroRequest, nil
}

// splitPONumbers transforms a comma-separated poNums string into []string
func splitPONumbers(poNums string) []string {
	if strings.TrimSpace(poNums) == "" {
		return nil
	}
	parts := strings.Split(poNums, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		t := strings.TrimSpace(p)
		if t != "" {
			out = append(out, t)
		}
	}
	return out
}
