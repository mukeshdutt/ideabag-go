package main

import (
	"fmt"
	"math"
)

type str struct {
	Name string
}

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

	// using openweather api mean by city name
	// or by latitude and longitude

	// latitude and longitude of delhi
	lat1 := 28.9875
	lon1 := 79.4141

	// latitude and longitude of hydrabad
	lat2 := 28.7041
	lon2 := 77.1025

	result := distace(lat1, lon1, lat2, lon2)
	fmt.Println("Result is: ", result)

	var input string
	for {

		println("hello")
		// in which unit you want to see
		// enter city 1
		// enter latitude and logitude for city 1 separated by comma
		// enter city 2
		// enter latitude and longigude fPor city 2 separted by comma
		// do you want to continue

		fmt.Printf("Please provide your input: ")
		fmt.Scanf("%s", &input)
		if input == "exit" {
			break
		} else {
			fmt.Println(input)
		}
	}
}
