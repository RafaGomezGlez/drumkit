package model

type Pagination struct {
	Start              int     `json:"start"`
	PageSize           int     `json:"pageSize"`
	TotalRecordsInPage int     `json:"totalRecordsInPage"`
	MoreAvailable      bool    `json:"moreAvailable"`
	LastObjectKey      *string `json:"lastObjectKey,omitempty"`
}

type StatusCode struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Status struct {
	Code StatusCode `json:"code"`
}

type ParentAccount struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   int    `json:"id"`
}

type Customer struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	ParentAccount ParentAccount `json:"parentAccount"`
}

type CustomerOrder struct {
	ID       int      `json:"id"`
	Customer Customer `json:"customer"`
	Deleted  bool     `json:"deleted"`
}

type Carrier struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	ParentAccount ParentAccount `json:"parentAccount"`
}

type CarrierOrder struct {
	ID      int     `json:"id"`
	Carrier Carrier `json:"carrier"`
	Deleted bool    `json:"deleted"`
}

type Shipment struct {
	ID            int             `json:"id"`
	CustomID      string          `json:"customId"`
	Status        Status          `json:"status"`
	CustomerOrder []CustomerOrder `json:"customerOrder"`
	CarrierOrder  []CarrierOrder  `json:"carrierOrder"`
	Created       string          `json:"created"`
	Updated       string          `json:"updated"`
	LastUpdatedOn string          `json:"lastUpdatedOn"`
	CreatedDate   string          `json:"createdDate"`
}

type ShipmentsDetails struct {
	Pagination Pagination `json:"pagination"`
	Shipments  []Shipment `json:"shipments"`
}

type ShipmentsResponse struct {
	Status  string           `json:"Status"`
	Details ShipmentsDetails `json:"details"`
}
