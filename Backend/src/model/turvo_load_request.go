package model

type AvroLoadRequest struct {
	LTLShipment             bool                `json:"ltlShipment"`
	StartDate               DateTime            `json:"startDate"`
	EndDate                 DateTime            `json:"endDate"`
	Status                  AvroStatus          `json:"status"`
	Groups                  []Group             `json:"groups"`
	Contributors            []Contributor       `json:"contributors"`
	Equipment               []Equipment         `json:"equipment"`
	Lane                    Lane                `json:"lane"`
	GlobalRoute             []GlobalRoute       `json:"globalRoute"`
	SkipDistanceCalculation bool                `json:"skipDistanceCalculation"`
	ModeInfo                []ModeInfo          `json:"modeInfo"`
	CustomerOrder           []CustomerOrderAvro `json:"customerOrder"`
	CarrierOrder            []CarrierOrderAvro  `json:"carrierOrder"`
	UseRoutingGuide         bool                `json:"use_routing_guide"`
}

type DateTime struct {
	Date     string `json:"date"`
	TimeZone string `json:"timeZone"`
}

type AvroStatus struct {
	Code        ValueKey `json:"code"`
	Notes       string   `json:"notes"`
	Description string   `json:"description"`
}

type ValueKey struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

type Group struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Operation int    `json:"_operation"`
}

type Contributor struct {
	Type            ValueKey        `json:"type"`
	ContributorUser ContributorUser `json:"contributorUser"`
	Operation       int             `json:"_operation"`
}

type ContributorUser struct {
	ID int `json:"id"`
}

type Equipment struct {
	Operation      int      `json:"_operation"`
	Type           ValueKey `json:"type"`
	Size           ValueKey `json:"size"`
	Weight         int      `json:"weight"`
	WeightUnits    ValueKey `json:"weightUnits"`
	Temp           int      `json:"temp"`
	TempUnits      ValueKey `json:"tempUnits"`
	ShipmentLength int      `json:"shipmentLength"`
}

type Lane struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type GlobalRoute struct {
	GlobalShipLocationSourceID string                     `json:"globalShipLocationSourceId"`
	Name                       string                     `json:"name"`
	SchedulingType             ValueKey                   `json:"schedulingType"`
	StopType                   ValueKey                   `json:"stopType"`
	Timezone                   string                     `json:"timezone"`
	Location                   Location                   `json:"location"`
	SegmentSequence            int                        `json:"segmentSequence"`
	LayoverTime                LayoverTime                `json:"layoverTime"`
	Sequence                   int                        `json:"sequence"`
	State                      string                     `json:"state"`
	Appointment                Appointment                `json:"appointment"`
	AppointmentConfirmation    bool                       `json:"appointmentConfirmation"`
	PlannedAppointmentDate     PlannedAppointmentDate     `json:"plannedAppointmentDate"`
	Services                   []ValueKey                 `json:"services"`
	PoNumbers                  []string                   `json:"poNumbers"`
	Notes                      string                     `json:"notes"`
	CustomerOrder              []GlobalRouteCustomerOrder `json:"customerOrder"`
	CarrierOrder               []GlobalRouteCarrierOrder  `json:"carrierOrder"`
	Transportation             Transportation             `json:"transportation"`
	FragmentDistance           Distance                   `json:"fragmentDistance"`
	Distance                   Distance                   `json:"distance"`
	StopLevelFragmentDistance  float64                    `json:"stop_level_fragment_distance,omitempty"`
}

type Location struct {
	ID int `json:"id"`
}

type LayoverTime struct {
	Value int      `json:"value"`
	Units ValueKey `json:"units"`
}

type Appointment struct {
	Date     string `json:"date"`
	Timezone string `json:"timezone"`
	Flex     int    `json:"flex"`
	HasTime  bool   `json:"hasTime"`
}

type PlannedAppointmentDate struct {
	SchedulingType ValueKey                 `json:"schedulingType"`
	Appointment    PlannedAppointmentWindow `json:"appointment"`
}

type PlannedAppointmentWindow struct {
	From Appointment `json:"from"`
	To   Appointment `json:"to"`
}

type GlobalRouteCustomerOrder struct {
	CustomerID            int `json:"customerId"`
	CustomerOrderSourceID int `json:"customerOrderSourceId"`
}

type GlobalRouteCarrierOrder struct {
	CarrierID            int `json:"carrierId"`
	CarrierOrderSourceID int `json:"carrierOrderSourceId"`
}

type Transportation struct {
	Mode        ValueKey `json:"mode"`
	ServiceType ValueKey `json:"serviceType"`
}

type Distance struct {
	Value float64  `json:"value"`
	Units ValueKey `json:"units"`
}

type ModeInfo struct {
	Operation             int               `json:"_operation"`
	SourceSegmentSequence string            `json:"sourceSegmentSequence"`
	Mode                  ValueKey          `json:"mode"`
	ServiceType           ValueKey          `json:"serviceType"`
	TotalSegmentValue     TotalSegmentValue `json:"totalSegmentValue"`
}

type TotalSegmentValue struct {
	Sync     bool     `json:"sync"`
	Value    int      `json:"value"`
	Currency ValueKey `json:"currency"`
}

type CustomerOrderAvro struct {
	CustomerOrderSourceID int          `json:"customerOrderSourceId"`
	Customer              CustomerAvro `json:"customer"`
	Items                 []Item       `json:"items"`
	Costs                 Costs        `json:"costs"`
	ExternalIds           []ExternalID `json:"externalIds"`
}

type CustomerAvro struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	Dimensions           Dimensions           `json:"dimensions"`
	ItemCategory         ValueKey             `json:"itemCategory"`
	Qty                  int                  `json:"qty"`
	Unit                 ValueKey             `json:"unit"`
	HandlingQty          int                  `json:"handlingQty"`
	HandlingUnit         ValueKey             `json:"handlingUnit"`
	Name                 string               `json:"name"`
	Notes                string               `json:"notes"`
	PickupLocation       []ItemLocation       `json:"pickupLocation"`
	DeliveryLocation     []ItemLocation       `json:"deliveryLocation"`
	Operation            int                  `json:"_operation"`
	ItemNumber           string               `json:"itemNumber"`
	Nmfc                 string               `json:"nmfc"`
	NmfcSub              string               `json:"nmfcSub"`
	IsHazmat             bool                 `json:"isHazmat"`
	Stackable            bool                 `json:"stackable"`
	FreightClass         ValueKey             `json:"freightClass"`
	Value                float64              `json:"value"`
	TotalValue           float64              `json:"totalValue"`
	Currency             ValueKey             `json:"currency"`
	MinTemp              Temperature          `json:"minTemp"`
	MaxTemp              Temperature          `json:"maxTemp"`
	StackDimensionsLimit StackDimensionsLimit `json:"stackDimensionsLimit"`
	LoadBearingCapacity  LoadBearingCapacity  `json:"loadBearingCapacity"`
	MaxStackCount        int                  `json:"maxStackCount"`
}

type Dimensions struct {
	Length int      `json:"length"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Units  ValueKey `json:"units"`
}

type ItemLocation struct {
	GlobalShipLocationSourceID string `json:"globalShipLocationSourceId"`
	Name                       string `json:"name"`
}

type Temperature struct {
	Temp     int      `json:"temp"`
	TempUnit ValueKey `json:"tempUnit"`
}

type StackDimensionsLimit struct {
	Height int      `json:"height"`
	Width  int      `json:"width"`
	Unit   ValueKey `json:"unit"`
}

type LoadBearingCapacity struct {
	Value int      `json:"value"`
	Unit  ValueKey `json:"unit"`
}

type Costs struct {
	TotalAmount float64    `json:"totalAmount"`
	LineItem    []LineItem `json:"lineItem"`
}

type LineItem struct {
	Code      ValueKey `json:"code"`
	Qty       int      `json:"qty"`
	Price     float64  `json:"price"`
	Amount    float64  `json:"amount"`
	Billable  bool     `json:"billable"`
	Notes     string   `json:"notes"`
	Operation int      `json:"_operation"`
}

type ExternalID struct {
	Type               ValueKey `json:"type"`
	Value              string   `json:"value"`
	CopyToCarrierOrder bool     `json:"copyToCarrierOrder"`
}

type CarrierOrderAvro struct {
	CarrierOrderSourceID int         `json:"carrierOrderSourceId"`
	Carrier              CarrierAvro `json:"carrier"`
	Drivers              []Driver    `json:"drivers"`
}

type CarrierAvro struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Driver struct {
	DriverID        int `json:"driverId"`
	Operation       int `json:"_operation"`
	SegmentSequence int `json:"segmentSequence"`
}
