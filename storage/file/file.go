package file

import "encoding/json"
import "os"
import "log"
import "io/ioutil"
import "strings"

import "trainstations_domain/stations"
import "trainstations_domain/lines"

type FileStationStorage struct {
	Stations []stations.Station `json:"Stations"`
	File     *os.File
}

type FileLineStorage struct {
	Lines []lines.Line `json:"Line"`
	File  *os.File
}

func NewLineStorage(file string) *FileLineStorage {
	storage := new(FileLineStorage)
	newFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	defer newFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	storage.File = newFile
	return storage
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

func (f *FileLineStorage) GetAll() ([]lines.Line, error) {
	jsonFile, err := os.Open(f.File.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)

	var data []lines.Line
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
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

func (f *FileStationStorage) Get(abbr string) ([]stations.Station, error) {
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

	for _, station := range data {
		if strings.ToLower(station.Abbr) == strings.ToLower(abbr) {
			selected := []stations.Station{station}
			return selected, nil
		}
	}

	return data, nil
}
