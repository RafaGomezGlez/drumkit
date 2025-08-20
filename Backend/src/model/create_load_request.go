package model

import "strings"

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

type ShipmentStatusCode string

const (
	StatusQuoteActive       ShipmentStatusCode = "2100"
	StatusTendered          ShipmentStatusCode = "2101"
	StatusCovered           ShipmentStatusCode = "2102"
	StatusDispatched        ShipmentStatusCode = "2103"
	StatusAtPickup          ShipmentStatusCode = "2104"
	StatusEnRoute           ShipmentStatusCode = "2105"
	StatusAtDelivery        ShipmentStatusCode = "2106"
	StatusDelivered         ShipmentStatusCode = "2107"
	StatusReadyForBilling   ShipmentStatusCode = "2108"
	StatusProcessing        ShipmentStatusCode = "2109"
	StatusCarrierPaid       ShipmentStatusCode = "2110"
	StatusCustomerPaid      ShipmentStatusCode = "2111"
	StatusCompleted         ShipmentStatusCode = "2112"
	StatusCanceled          ShipmentStatusCode = "2113"
	StatusQuoteInactive     ShipmentStatusCode = "2114"
	StatusPickedUp          ShipmentStatusCode = "2115"
	StatusRouteComplete     ShipmentStatusCode = "2116"
	StatusTenderOffered     ShipmentStatusCode = "2117"
	StatusTenderAccepted    ShipmentStatusCode = "2118"
	StatusTenderRejected    ShipmentStatusCode = "2119"
	StatusDraft             ShipmentStatusCode = "2120"
	StatusShipmentReady     ShipmentStatusCode = "2121"
	StatusAcquiringLocation ShipmentStatusCode = "2123"
	StatusCustomsHold       ShipmentStatusCode = "2124"
	StatusArrived           ShipmentStatusCode = "2125"
	StatusAvailable         ShipmentStatusCode = "2126"
	StatusOutGated          ShipmentStatusCode = "2127"
	StatusInGated           ShipmentStatusCode = "2129"
	StatusArrivingToPort    ShipmentStatusCode = "2131"
	StatusBerthing          ShipmentStatusCode = "2132"
	StatusUnloading         ShipmentStatusCode = "2133"
	StatusRamped            ShipmentStatusCode = "2134"
	StatusDeramped          ShipmentStatusCode = "2135"
	StatusDeparted          ShipmentStatusCode = "2136"
	StatusHeld              ShipmentStatusCode = "2137"
	StatusOutForDelivery    ShipmentStatusCode = "2138"
	StatusInTransShipment   ShipmentStatusCode = "2139"
	StatusOnHold            ShipmentStatusCode = "2140"
)

// StatusCodeForValue converts a human-friendly status value (e.g. "Covered", "Quote")
// into the numeric code string. Returns empty string and false if unknown.
func StatusCodeForValue(value string) (ShipmentStatusCode, bool) {
	if value == "" {
		return "", false
	}
	key := strings.ToLower(strings.TrimSpace(value))

	var m = map[string]ShipmentStatusCode{
		"quote":              StatusQuoteActive,
		"quote active":       StatusQuoteActive,
		"tendered":           StatusTendered,
		"covered":            StatusCovered,
		"dispatched":         StatusDispatched,
		"at pickup":          StatusAtPickup,
		"en route":           StatusEnRoute,
		"at delivery":        StatusAtDelivery,
		"delivered":          StatusDelivered,
		"ready for billing":  StatusReadyForBilling,
		"processing":         StatusProcessing,
		"carrier paid":       StatusCarrierPaid,
		"customer paid":      StatusCustomerPaid,
		"completed":          StatusCompleted,
		"canceled":           StatusCanceled,
		"quote inactive":     StatusQuoteInactive,
		"picked up":          StatusPickedUp,
		"route complete":     StatusRouteComplete,
		"tender offered":     StatusTenderOffered,
		"tender accepted":    StatusTenderAccepted,
		"tender rejected":    StatusTenderRejected,
		"draft":              StatusDraft,
		"shipment ready":     StatusShipmentReady,
		"acquiring location": StatusAcquiringLocation,
		"customs hold":       StatusCustomsHold,
		"arrived":            StatusArrived,
		"available":          StatusAvailable,
		"out gated":          StatusOutGated,
		"in gated":           StatusInGated,
		"arriving to port":   StatusArrivingToPort,
		"berthing":           StatusBerthing,
		"unloading":          StatusUnloading,
		"ramped":             StatusRamped,
		"deramped":           StatusDeramped,
		"departed":           StatusDeparted,
		"held":               StatusHeld,
		"out for delivery":   StatusOutForDelivery,
		"in trans shipment":  StatusInTransShipment,
		"on hold":            StatusOnHold,
	}

	// try exact match; also try simplified key without extra words
	if code, ok := m[key]; ok {
		return code, true
	}
	// fallback: try first word (e.g., "Quote" from "Quote - something")
	parts := strings.Fields(key)
	if len(parts) > 0 {
		if code, ok := m[parts[0]]; ok {
			return code, true
		}
	}
	return "", false
}
