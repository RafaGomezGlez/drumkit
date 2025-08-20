package customer

// Models for Turvo customers list response

type CustomersResponse struct {
	Status  string           `json:"Status"`
	Details CustomersDetails `json:"details"`
}

type CustomersDetails struct {
	Pagination Pagination       `json:"pagination"`
	Customers  []Customer       `json:"customers"`
	Billings   []BillingsLookup `json:"billings,omitempty"`
}

type Pagination struct {
	Start              int     `json:"start"`
	PageSize           int     `json:"pageSize"`
	TotalRecordsInPage int     `json:"totalRecordsInPage"`
	MoreAvailable      bool    `json:"moreAvailable"`
	LastObjectKey      *string `json:"lastObjectKey"`
}

type Customer struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Billing       BillingSummary `json:"billing"`
	LastUpdatedOn string         `json:"lastUpdatedOn"`
	Updated       string         `json:"updated"`
	CreatedDate   string         `json:"createdDate"`
	Created       string         `json:"created"`
	Owner         Owner          `json:"owner"`
	Status        Status         `json:"status"`
	Billings      []BillingEntry `json:"billings,omitempty"`
	Address       []AddressEntry `json:"address,omitempty"`
	ExternalIds   []ExternalID   `json:"externalIds,omitempty"`
	Email         []EmailEntry   `json:"email,omitempty"`
	Phone         []PhoneEntry   `json:"phone,omitempty"`
}

type Owner struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Status struct {
	Code        int    `json:"code"`
	Notes       string `json:"notes"`
	Description string `json:"description"`
}

type BillingSummary struct {
	ToName string `json:"toName"`
	// optional fields
	PayTerms    string  `json:"payTerms,omitempty"`
	CreditLimit float64 `json:"creditLimit,omitempty"`
}

type BillingEntry struct {
	ID           string        `json:"id"`
	ThirdParty   bool          `json:"thirdParty"`
	ToName       string        `json:"toName"`
	AutoInvoice  bool          `json:"autoInvoice"`
	Deleted      bool          `json:"deleted"`
	CreditLimit  float64       `json:"creditLimit,omitempty"`
	Instructions *string       `json:"instructions,omitempty"`
	PayTerms     *ValueKey     `json:"payTerms,omitempty"`
	Address      NestedAddress `json:"address"`
}

type NestedAddress struct {
	ID      string  `json:"id"`
	Line1   string  `json:"line1,omitempty"`
	Line2   *string `json:"line2,omitempty"`
	Zip     string  `json:"zip,omitempty"`
	Country string  `json:"country,omitempty"`
	City    string  `json:"city,omitempty"`
	Lon     float64 `json:"lon,omitempty"`
	Lat     float64 `json:"lat,omitempty"`
}

type AddressEntry struct {
	ID        string    `json:"id"`
	Line1     string    `json:"line1,omitempty"`
	Line2     *string   `json:"line2,omitempty"`
	City      string    `json:"city,omitempty"`
	State     string    `json:"state,omitempty"`
	Country   string    `json:"country,omitempty"`
	Zip       string    `json:"zip,omitempty"`
	Type      *ValueKey `json:"type,omitempty"`
	IsPrimary bool      `json:"isPrimary,omitempty"`
	Deleted   bool      `json:"deleted,omitempty"`
}

type ValueKey struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type ExternalID struct {
	ID      string    `json:"id"`
	Type    *ValueKey `json:"type,omitempty"`
	Value   string    `json:"value,omitempty"`
	Deleted bool      `json:"deleted,omitempty"`
}

type EmailEntry struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Deleted bool      `json:"deleted,omitempty"`
	Type    *ValueKey `json:"type,omitempty"`
}

type PhoneEntry struct {
	ID        string    `json:"id"`
	Number    string    `json:"number"`
	Extension string    `json:"extension,omitempty"`
	Primary   bool      `json:"primary,omitempty"`
	Deleted   bool      `json:"deleted,omitempty"`
	Type      *ValueKey `json:"type,omitempty"`
	Country   *ValueKey `json:"country,omitempty"`
}

// BillingsLookup models the odd "billings" summary object that contains address state lists
type BillingsLookup struct {
	Address struct {
		State []string `json:"state,omitempty"`
	} `json:"address,omitempty"`
}
