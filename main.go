package main

import "fmt"

import "trainstations_domain/stations"

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
	storageFile := "test.txt"

	stationRepository := file.NewFileStationStorage(storageFile)
	//fmt.Println(stationRepository)

	_ = stationRepository.Add(stations.Station{Name: "Montgomery", Abbr: "MONT"})
	_ = stationRepository.Add(stations.Station{Name: "North Concord", Abbr: "NCON"})

	allStations, _ := stationRepository.GetAll()
	fmt.Println(allStations)
}
