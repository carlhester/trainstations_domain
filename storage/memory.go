package storage

import "trainstations_domain/stations"

type MemoryStationStorage struct {
	stations []stations.Station
}

func NewMemoryStationStorage() *MemoryStationStorage {
	storage := stations.Station{Abbr: "MONT", Name: "Montgomery"}
	data := new(MemoryStationStorage)
	data.Add(storage)
	return data
}

// GetAllStations returns a slice of Station structs
func (m *MemoryStationStorage) GetAllStations() ([]stations.Station, error) {
	return m.stations, nil
}

func (m *MemoryStationStorage) Add(station stations.Station) error {
	m.stations = append(m.stations, station)
	return nil
}

func (m *MemoryStationStorage) GetStationByAbbr(abbr string) (stations.Station, error) {
	StationInMemory := stations.Station{Abbr: "MONT", Name: "Montgomery"}
	queriedStation := StationInMemory
	return queriedStation, nil
}
