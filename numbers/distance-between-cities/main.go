package main

import (
	"fmt"
	"math"
)

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func distace(lat1, lon1, lat2, lon2 float64) float64 {
	R := 6371
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := float64(R) * c
	return d
}

func main() {

	var input string
	for {

		/* --- Delhi's latitude and longitude
		lat1 := 28.9875
		lon1 := `79.4141

		--- Hydrabad latitude and longitude
		lat2 := 28.7041
		lon2 := 77.1025
		*/

		srcCity, destCity := "", ""
		fmt.Printf("Enter source city name:")
		fmt.Scanf("%s", &srcCity)
		fmt.Printf("Enter destination city name:")
		fmt.Scanf("%s", &destCity)

		srcLat, srcLon := 0.0, 0.0
		fmt.Print("Enter Source Latitude: ")
		fmt.Scanf("%f", &srcLat)
		fmt.Print("Enter Source Longitude: ")
		fmt.Scanf("%f", &srcLon)

		destLat, destLon := 0.0, 0.0
		fmt.Print("Enter Destination Latitude: ")
		fmt.Scanf("%f", &destLat)
		fmt.Print("Enter Destination Longitude:")
		fmt.Scanf("%f", &destLon)

		distance := distace(srcLat, srcLon, destLat, destLon)
		fmt.Printf("Distance between two cities %s and %s is: %f \n", srcCity, destCity, distance)

		// do you want to continue
		fmt.Print("Do you want to continue, press n to exit and any other keys to continue:")
		fmt.Scanf("%s", &input)
		if input == "n" || input == "no" {
			break
		} else {
			fmt.Println(input)
		}
	}
}
