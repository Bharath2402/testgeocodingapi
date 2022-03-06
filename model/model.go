package model

/* This model can be extended based on required asserions
The Locality, Country and Postal Code are considered here for simplicity*/
type GeoCodeTestPositive struct {
	TestDesc        string
	Lattitude       string
	Longtitude      string
	StatusCode      int
	RespFullAddress string
	RespLocality    string
	RespCountry     string
	RespPinCode     string
	RespBodyStatus  string
}

type QueryString struct {
	QueryKey   string
	QueryValue string
}

type GeoCodeTestNegative struct {
	TestDesc   string
	Lattitude  string
	Longtitude string
	StatusCode int
	RespStatus string
	RespError  string
}

type GeoCodeTestNegativeApiKey struct {
	TestDesc   string
	Lattitude  string
	Longtitude string
	ApiKey     string
	StatusCode int
	RespStatus string
	RespError  string
}
