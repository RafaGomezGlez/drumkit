package locations

import "time"

type LocationsResponse struct {
	Status  string          `json:"Status"`
	Details LocationsDetail `json:"details"`
}

type LocationsDetail struct {
	Pagination Pagination `json:"pagination"`
	Locations  []Location `json:"locations"`
}

type Pagination struct {
	Start              int  `json:"start"`
	PageSize           int  `json:"pageSize"`
	TotalRecordsInPage int  `json:"totalRecordsInPage"`
	MoreAvailable      bool `json:"moreAvailable"`
}

type Location struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Addresses []Address `json:"addresses"`
	Phones    []Phone   `json:"phones"`
}

type Address struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	City  string `json:"city"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}

type Phone struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
	Extension   string `json:"extension"`
}
