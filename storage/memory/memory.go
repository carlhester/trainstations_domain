package memory

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

func GetStationByAbbr(abbr string) (stations.Station, error) {
	_ = abbr
	Station := stations.Station{Abbr: "MONT", Name: "Montgomery"}
	return Station, nil
}
