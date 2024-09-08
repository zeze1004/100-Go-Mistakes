package main

import (
	"fmt"
	"log"
)

type Route struct{}

func GetRoute1(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates1(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates") // 에러 로깅하고 return
		return Route{}, err
	}

	err = validateCoordinates1(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates") // 에러 로깅하고 return
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates1(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat)
		return fmt.Errorf("invalid latitude: %f", lat) // 에러 로깅하고 return
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng)
		return fmt.Errorf("invalid longitude: %f", lng) // 에러 로깅하고 return
	}
	return nil
}

func getRoute(lat, lng, lat2, lng2 float32) (Route, error) {
	return Route{}, nil
}
