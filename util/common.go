package util

const commaStr = ","

func MakeLatlngString(lat string, lng string) (string, error) {

	return lat + commaStr + lng, nil
}
