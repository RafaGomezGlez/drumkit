package model

// CreateLoadRequest is the top-level struct for the create load request body.
type CreateLoadRequest struct {
	ExternalTMSLoadID string           `json:"externalTMSLoadID"`
	FreightLoadID     string           `json:"freightLoadID"`
	Status            string           `json:"status"`
	Customer          CLCustomer       `json:"customer"`
	BillTo            CLBillTo         `json:"billTo"`
	Pickup            CLPickup         `json:"pickup"`
	Consignee         CLConsignee      `json:"consignee"`
	Carrier           CLCarrier        `json:"carrier"`
	RateData          CLRateData       `json:"rateData"`
	Specifications    CLSpecifications `json:"specifications"`
	InPalletCount     int              `json:"inPalletCount"`
	OutPalletCount    int              `json:"outPalletCount"`
	NumCommodities    int              `json:"numCommodities"`
	TotalWeight       float64          `json:"totalWeight"`
	BillableWeight    float64          `json:"billableWeight"`
	PoNums            string           `json:"poNums"`
	Operator          string           `json:"operator"`
	RouteMiles        float64          `json:"routeMiles"`
}

// CLCustomer represents the customer object in the request.
// Using a "CL" prefix to avoid name collision with the existing Customer struct in shipment.go.
type CLCustomer struct {
	ExternalTMSId string `json:"externalTMSId"`
	Name          string `json:"name"`
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
	Country       string `json:"country"`
	Contact       string `json:"contact"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	RefNumber     string `json:"refNumber"`
}

// CLBillTo represents the billTo object.
type CLBillTo struct {
	ExternalTMSId string `json:"externalTMSId"`
	Name          string `json:"name"`
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
	Country       string `json:"country"`
	Contact       string `json:"contact"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
}

// CLPickup represents the pickup object.
type CLPickup struct {
	ExternalTMSId string `json:"externalTMSId"`
	Name          string `json:"name"`
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
	Country       string `json:"country"`
	Contact       string `json:"contact"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	BusinessHours string `json:"businessHours"`
	RefNumber     string `json:"refNumber"`
	ReadyTime     string `json:"readyTime"`
	ApptTime      string `json:"apptTime"`
	ApptNote      string `json:"apptNote"`
	Timezone      string `json:"timezone"`
	WarehouseId   string `json:"warehouseId"`
}

// CLConsignee represents the consignee object.
type CLConsignee struct {
	ExternalTMSId string `json:"externalTMSId"`
	Name          string `json:"name"`
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
	Country       string `json:"country"`
	Contact       string `json:"contact"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	BusinessHours string `json:"businessHours"`
	RefNumber     string `json:"refNumber"`
	MustDeliver   string `json:"mustDeliver"`
	ApptTime      string `json:"apptTime"`
	ApptNote      string `json:"apptNote"`
	Timezone      string `json:"timezone"`
	WarehouseId   string `json:"warehouseId"`
}

// CLCarrier represents the carrier object.
type CLCarrier struct {
	McNumber                 string `json:"mcNumber"`
	DotNumber                string `json:"dotNumber"`
	Name                     string `json:"name"`
	Phone                    string `json:"phone"`
	Dispatcher               string `json:"dispatcher"`
	SealNumber               string `json:"sealNumber"`
	Scac                     string `json:"scac"`
	FirstDriverName          string `json:"firstDriverName"`
	FirstDriverPhone         string `json:"firstDriverPhone"`
	SecondDriverName         string `json:"secondDriverName"`
	SecondDriverPhone        string `json:"secondDriverPhone"`
	Email                    string `json:"email"`
	DispatchCity             string `json:"dispatchCity"`
	DispatchState            string `json:"dispatchState"`
	ExternalTMSTruckId       string `json:"externalTMSTruckId"`
	ExternalTMSTrailerId     string `json:"externalTMSTrailerId"`
	ConfirmationSentTime     string `json:"confirmationSentTime"`
	ConfirmationReceivedTime string `json:"confirmationReceivedTime"`
	DispatchedTime           string `json:"dispatchedTime"`
	ExpectedPickupTime       string `json:"expectedPickupTime"`
	PickupStart              string `json:"pickupStart"`
	PickupEnd                string `json:"pickupEnd"`
	ExpectedDeliveryTime     string `json:"expectedDeliveryTime"`
	DeliveryStart            string `json:"deliveryStart"`
	DeliveryEnd              string `json:"deliveryEnd"`
	SignedBy                 string `json:"signedBy"`
	ExternalTMSId            string `json:"externalTMSId"`
}

// CLRateData represents the rateData object.
type CLRateData struct {
	CustomerRateType  string  `json:"customerRateType"`
	CustomerNumHours  float64 `json:"customerNumHours"`
	CustomerLhRateUsd float64 `json:"customerLhRateUsd"`
	FscPercent        float64 `json:"fscPercent"`
	FscPerMile        float64 `json:"fscPerMile"`
	CarrierRateType   string  `json:"carrierRateType"`
	CarrierNumHours   float64 `json:"carrierNumHours"`
	CarrierLhRateUsd  float64 `json:"carrierLhRateUsd"`
	CarrierMaxRate    float64 `json:"carrierMaxRate"`
	NetProfitUsd      float64 `json:"netProfitUsd"`
	ProfitPercent     float64 `json:"profitPercent"`
}

// CLSpecifications represents the specifications object.
type CLSpecifications struct {
	MinTempFahrenheit int  `json:"minTempFahrenheit"`
	MaxTempFahrenheit int  `json:"maxTempFahrenheit"`
	LiftgatePickup    bool `json:"liftgatePickup"`
	LiftgateDelivery  bool `json:"liftgateDelivery"`
	InsidePickup      bool `json:"insidePickup"`
	InsideDelivery    bool `json:"insideDelivery"`
	Tarps             bool `json:"tarps"`
	Oversized         bool `json:"oversized"`
	Hazmat            bool `json:"hazmat"`
	Straps            bool `json:"straps"`
	Permits           bool `json:"permits"`
	Escorts           bool `json:"escorts"`
	Seal              bool `json:"seal"`
	CustomBonded      bool `json:"customBonded"`
	Labor             bool `json:"labor"`
}
