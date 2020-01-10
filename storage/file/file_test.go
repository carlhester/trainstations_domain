package file_test

import "trainstations_domain/stations"
import "trainstations_domain/storage/file"
import "testing"

func TestAdd(t *testing.T) {
	t.Parallel()

	t.Run("add a known good station", func(t *testing.T) {
		stationRepository := new(file.FileStationStorage)
		have := stationRepository.Add(stations.Station{Abbr: "NCON", Name: "North Concord"})
		if have != nil {
			t.Fatal(have)
		}
	})

	t.Run("blank", func(t *testing.T) {
		stationRepository := new(file.FileStationStorage)
		have := stationRepository.Add(stations.Station{Abbr: "", Name: ""})
		if have != nil {
			t.Fatal(have)
		}
	})

	t.Run("add a dupe", func(t *testing.T) {
		stationRepository := new(file.FileStationStorage)
		_ = stationRepository.Add(stations.Station{Abbr: "MONT", Name: "Montgomery"})
		have := stationRepository.Add(stations.Station{Abbr: "MONT", Name: "Montgomery"})

		if have == nil {
			t.Fatal(have)
		}
	})

}
