package main

import "fmt"

import "trainstations_domain/stations"

//import "trainstations_domain/storage/memory"
import "trainstations_domain/storage/file"

func main() {

	// memory
	//stationRepository := new(memory.MemoryStationStorage)
	//allStations, _ := stationRepository.GetAll()
	//fmt.Println(allStations)

	//stationRepository.Add(stations.Station{Abbr: "MONT", Name: "Montgomery"})
	//allStations, _ = stationRepository.GetAll()
	//fmt.Println(allStations)

	//stationRepository.Add(stations.Station{Abbr: "NCON", Name: "North Concord"})
	//allStations, _ = stationRepository.GetAll()
	//fmt.Println(allStations)

	//mont, _ := stationRepository.Get("MONT")
	//fmt.Println(mont)

	// file
	storageFile := "test.txt"
	stationRepository := file.NewStationStorage(storageFile)

	_ = stationRepository.Add(stations.Station{Name: "Montgomery St.", Abbr: "MONT"})

	allStations, _ := stationRepository.GetAll()
	fmt.Println(allStations)
}
