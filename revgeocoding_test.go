package testgeocodingapi

import (
	"fmt"
	"log"
	"net/http"
	"testgeocodingapi/model"
	"testgeocodingapi/util"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestRevGeoCodingPositive(t *testing.T) {
	e := httpexpect.New(t, host)
	// TDT loop
	for i := 0; i < len(config.GeoCodePostiveCases); i++ {

		testObject := config.GeoCodePostiveCases[i]
		fmt.Println("Running positive scenario -> " + testObject.TestDesc)

		// Make latlng query
		latlngValue, err := util.MakeLatlngString(testObject.Lattitude, testObject.Longtitude)
		if err != nil {
			t.Error(err)
			log.Fatalf("Error in lattitude/longtitude convertion")
		}

		fullResp := e.GET("/api/geocode/json").
			WithQuery("key", apiKey).
			WithQuery("latlng", latlngValue).
			Expect().
			Status(http.StatusOK).JSON().Object()
		respAddress := fullResp.Value("results").Array().Element(0).Object().Value("address_components")

		// Assert address components

		respAddress.Array().Element(model.AddressComponentLocality).Object().
			ValueEqual("long_name", testObject.RespLocality)

		respAddress.Array().Element(model.AddressComponentCountry).Object().
			ValueEqual("long_name", testObject.RespCountry)

		respAddress.Array().Element(model.AddressComponentPinCode).Object().
			ValueEqual("long_name", testObject.RespPinCode)

		// Assert full formatted address

		fullResp.Value("results").Array().Element(0).Object().
			ValueEqual("formatted_address", testObject.RespFullAddress)
	}

}

func TestRevGeoCodingNegativeCases(t *testing.T) {
	e := httpexpect.New(t, host)
	// Invalid latlag scenarios
	for i := 0; i < len(config.GeoCodeNegativeCases); i++ {

		testObject := config.GeoCodeNegativeCases[i]
		fmt.Println("Running negative scenario -> " + testObject.TestDesc)

		// Make latlng query
		latlngValue, err := util.MakeLatlngString(testObject.Lattitude, testObject.Longtitude)
		if err != nil {
			t.Error(err)
			log.Fatalf("Error in lattitude/longtitude convertion")
		}

		resp := e.GET("/api/geocode/json").
			WithQuery("key", apiKey).
			WithQuery("latlng", latlngValue).
			Expect().
			Status(testObject.StatusCode).JSON().Object()
		resp.ValueEqual("status", testObject.RespStatus)
		resp.ValueEqual("error_message", testObject.RespError)
	}

	// invalid api key scenarios
	for i := 0; i < len(config.GeoCodeNegativekeyCases); i++ {

		testObject := config.GeoCodeNegativekeyCases[i]
		fmt.Println("Running negative scenario -> " + testObject.TestDesc)
		// Make latlng query
		latlngValue, err := util.MakeLatlngString(testObject.Lattitude, testObject.Longtitude)
		if err != nil {
			t.Error(err)
			log.Fatalf("Error in lattitude/longtitude convertion")
		}

		resp := e.GET("/api/geocode/json").
			WithQuery("key", testObject.ApiKey).
			WithQuery("latlng", latlngValue).
			Expect().
			Status(testObject.StatusCode).JSON().Object()
		resp.ValueEqual("status", testObject.RespStatus)
		resp.ValueEqual("error_message", testObject.RespError)
	}

}
