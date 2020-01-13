package memory

import "errors"
import "trainstations_domain/stations"

type MemoryStationStorage struct {
	stations []stations.Station
}

// This constructor exists because when we're storing in memory we want to
// initialize our empty MemoryStationStorage with some data.
func NewStationStorage() *MemoryStationStorage {
	storage := stations.Station{Abbr: "MONT", Name: "Montgomery"}
	data := new(MemoryStationStorage)
	data.Add(storage)
	return data
}

// Adds a new Station to to the slice of structs
func (m *MemoryStationStorage) Add(station stations.Station) error {
	for _, existingStation := range m.stations {
		if existingStation.Name == station.Name &&
			existingStation.Abbr == station.Abbr {
			return errors.New("duplicate")
		}
	}
	m.stations = append(m.stations, station)
	return nil
}

// Returns a Station when provided with an existing Abbr
// Otherwise returns an empty Station and error
func (m *MemoryStationStorage) Get(abbr string) ([]stations.Station, error) {
	for _, station := range m.stations {
		if station.Abbr == abbr {
			selected := []stations.Station{station}
			return selected, nil
		}
	}
	return m.stations, nil
}

func (m *MemoryStationStorage) GetAll() ([]stations.Station, error) {
	return m.stations, nil
}
