package nominatim

type Address struct {
	Road        string `json:"road"`
	City        string `json:"city"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}
