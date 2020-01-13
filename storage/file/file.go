package file

import "encoding/json"
import "os"
import "log"
import "io/ioutil"

import "trainstations_domain/stations"

type FileStationStorage struct {
	Stations []stations.Station `json:"Stations"`
	File     *os.File
}

func NewStationStorage(file string) *FileStationStorage {
	storage := new(FileStationStorage)
	newFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	storage.File = newFile
	return storage
}

func (f *FileStationStorage) Add(submittedStation stations.Station) error {
	//newStation, _ := json.Marshal(station)
	oldStations, _ := f.GetAll()
	var allStations []stations.Station

	for _, eachStation := range oldStations {
		allStations = append(allStations, eachStation)
	}

	for _, everyStation := range allStations {
		if submittedStation == everyStation {
			return nil
		}
	}

	allStations = append(allStations, submittedStation)

	d, _ := json.Marshal(allStations)

	err := ioutil.WriteFile(f.File.Name(), d, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (f *FileStationStorage) GetAll() ([]stations.Station, error) {
	jsonFile, err := os.Open(f.File.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)

	var data []stations.Station
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
