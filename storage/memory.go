package storage

import "errors"
import "trainstations_domain/stations"

type MemoryStationStorage struct {
	stations []stations.Station
}

// This constructor exists because when we're storing in memory we want to
// initialize our empty MemoryStationStorage with some data.
func NewMemoryStationStorage() *MemoryStationStorage {
	storage := stations.Station{Abbr: "MONT", Name: "Montgomery"}
	data := new(MemoryStationStorage)
	data.Add(storage)
	return data
}

// GetAllStations returns a slice of all the Station structs
func (m *MemoryStationStorage) GetAllStations() ([]stations.Station, error) {
	return m.stations, nil
}

// Adds a new Station to to the slice of structs
func (m *MemoryStationStorage) Add(station stations.Station) error {
	m.stations = append(m.stations, station)
	return nil
}

// Returns a single Station when passed a valid Station.Abbr
func (m *MemoryStationStorage) GetStationByAbbr(abbr string) (stations.Station, error) {
	stationInMemory := []stations.Station{{Abbr: "MONT", Name: "Montgomery"}}
	for _, station := range stationInMemory {
		if station.Abbr == abbr {
			return station, nil
		}
	}
	blankStation := stations.Station{}
	return blankStation, errors.New("fail")
}
