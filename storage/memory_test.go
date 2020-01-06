package storage

import "trainstations_domain/stations"
import "testing"

func TestGetAllStations(t *testing.T) {
	stationRepository := NewMemoryStationStorage()

	have, _ := stationRepository.GetAllStations()
	want := []stations.Station{{Abbr: "MONT", Name: "Montgomery"}}

	if (have[0].Name != want[0].Name) || (have[0].Abbr != want[0].Abbr) {
		t.Error(have[0], want[0])
	}
}

func TestGetStationByAbbr(t *testing.T) {
	stationRepository := NewMemoryStationStorage()

	have, _ := stationRepository.GetStationByAbbr("MONT")
	want := stations.Station{Abbr: "MONT", Name: "Montgomery"}

	if have != want {
		t.Error(have, want)
	}
}

func TestGetStationByAbbrFailLookup(t *testing.T) {
	stationRepository := NewMemoryStationStorage()

	have, err := stationRepository.GetStationByAbbr("")

	if err == nil {
		t.Error(have, err)
	}
}
