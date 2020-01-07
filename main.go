package main

import "fmt"

import "trainstations_domain/stations"
import "trainstations_domain/storage"

func main() {

	stationRepository := new(storage.MemoryStationStorage)
	allStations, _ := stationRepository.GetAll()
	fmt.Println(allStations)

	stationRepository.Add(stations.Station{Abbr: "MONT", Name: "Montgomery"})
	allStations, _ = stationRepository.GetAll()
	fmt.Println(allStations)

	stationRepository.Add(stations.Station{Abbr: "NCON", Name: "North Concord"})
	allStations, _ = stationRepository.GetAll()
	fmt.Println(allStations)

	mont, _ := stationRepository.Get("MONT")
	fmt.Println(mont)
}
