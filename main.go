package main

import "fmt"

//import "trainstations_domain/stations"
//import "trainstations_domain/storage/memory"
import "trainstations_domain/storage/file"

func main() {

	// memory
	//	stationRepository := new(memory.MemoryStationStorage)
	//	allStations, _ := stationRepository.GetAll()
	//	fmt.Println(allStations)
	//
	//	stationRepository.Add(stations.Station{Abbr: "MONT", Name: "Montgomery"})
	//	allStations, _ = stationRepository.GetAll()
	//	fmt.Println(allStations)
	//
	//	stationRepository.Add(stations.Station{Abbr: "NCON", Name: "North Concord"})
	//	allStations, _ = stationRepository.GetAll()
	//	fmt.Println(allStations)
	//
	//	mont, _ := stationRepository.Get("MONT")
	//	fmt.Println(mont)
	//
	// file
	stationRepository := file.NewFileStationStorage()
	fmt.Println(stationRepository)

	//allStations, _ = stationRepository.GetAll()
	//fmt.Println(allStations)
}
