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

// Adds a new Station to to the slice of structs
func (m *MemoryStationStorage) Add(station stations.Station) error {
	m.stations = append(m.stations, station)
	return nil
}

func (m *MemoryStationStorage) Get(abbr string) (stations.Station, error) {
	for _, station := range m.stations {
		if station.Abbr == abbr {
			return station, nil
		}
	}
	emptyStation := stations.Station{}
	return emptyStation, errors.New("fail")
}

func (m *MemoryStationStorage) GetAll() ([]stations.Station, error) {
	return m.stations, nil
}
