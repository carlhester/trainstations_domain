package main

import "testing"
import "trainstations_domain/stations"
import "trainstations_domain/storage/memory"

func TestAddThenRetrieve(t *testing.T) {
	t.Parallel()
	repo := memory.NewStationStorage()
	stationToAdd := stations.Station{Name: "North Concord", Abbr: "NCON"}
	_ = repo.Add(stationToAdd)
	resp, _ := repo.GetByAbbr("NCON")
	if resp != stationToAdd {
		t.Fail()
	}

}
