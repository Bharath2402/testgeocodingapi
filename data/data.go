package data

import (
	"log"

	"github.com/pelletier/go-toml"
)

const testConf = "conf/init.toml"

var Url, GeoApikey string

func init() {
	var err error

	config, err := toml.LoadFile(testConf)
	if err != nil {
		log.Fatal("Error while loading test data ", err)
	}

	Url = config.Get("init.url").(string)
	GeoApikey = config.Get("init.geoApiKey").(string)
}
