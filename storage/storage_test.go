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
		t.Fatal(have, want)
	}
}

func TestGetStationByAbbrFailLookup(t *testing.T) {
	stationRepository := NewMemoryStationStorage()

	t.Run("blank", func(t *testing.T) {
		have, err := stationRepository.GetStationByAbbr("")
		if err == nil {
			t.Fatal(have, err)
		}
	})

	t.Run("noisy", func(t *testing.T) {
		have, err := stationRepository.GetStationByAbbr("h4oih2ohgo4h2o4h2o4h")
		if err == nil {
			t.Fatal(have, err)
		}
	})
}
