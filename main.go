package main

import "fmt"
import "trainstations_domain/stations"
import "trainstations_domain/storage"

func main() {

	// Create a variable to serve domain functionality
	var stationRepository stations.Repository

	// Set our repository to use the in-memory storage
	stationRepository = storage.NewMemoryStationStorage()

	// Use the GetAllStations functionality to populate stations
	stations, _ := stationRepository.GetAllStations()

	// Print result
	fmt.Println(stations)
}
