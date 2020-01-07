package storage

import "trainstations_domain/stations"

type FileStationStorage struct {
	stations []stations.Station
}

// GetAllStations returns a slice of all the Station structs
func (m *FileStationStorage) GetAllStations() ([]stations.Station, error) {
	return m.stations, nil
}
