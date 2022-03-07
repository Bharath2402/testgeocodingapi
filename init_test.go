package testgeocodingapi

import (
	"io/ioutil"
	"log"
	"strings"
	"testgeocodingapi/data"
	"testgeocodingapi/model"
	"testing"

	"github.com/pelletier/go-toml"
)

var host, apiKey string

// Declare test data config

type Config struct {
	GeoCodePostiveCases     []model.GeoCodeTestPositive
	GeoCodeNegativeCases    []model.GeoCodeTestNegative
	GeoCodeNegativekeyCases []model.GeoCodeTestNegativeApiKey
}

var config = Config{}

// Starting point of test execution

func TestMain(m *testing.M) {

	// 1. Test setup - optional
	host = data.Url
	apiKey = data.GeoApikey

	if len(strings.TrimSpace(apiKey)) == 0 {
		log.Fatalf("Please enter a  API key and run the tests")
	}

	configTomlAsByte, err := ioutil.ReadFile("conf/testdata.toml")
	if err != nil {
		log.Fatal(err)
	}
	toml.Unmarshal(configTomlAsByte, &config)

	// 2. Run tests
	m.Run()

	// 3. Test teardown - optional
}
