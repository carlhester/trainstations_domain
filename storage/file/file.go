package file

import "encoding/json"
import "os"
import "log"
import "fmt"
import "io/ioutil"

import "trainstations_domain/stations"

type FileStationStorage struct {
	Stations []stations.Station
	File     *os.File
}

func NewFileStationStorage(file string) *FileStationStorage {
	storage := new(FileStationStorage)
	newFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	storage.File = newFile
	return storage

}

func (f *FileStationStorage) Add(station stations.Station) error {
	b, _ := json.Marshal(station)
	bytes := []byte(string(b))
	bytesWritten, err := f.File.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
	_ = f.File.Sync()
	return nil
}

func (f *FileStationStorage) GetAll() ([]stations.Station, error) {
	data, err := ioutil.ReadFile(f.File.Name())
	if err != nil {
		log.Fatal(err)
	}

	var data2 stations.Station

	err = json.Unmarshal(data, &data2)
	fmt.Println(data2)

	return f.Stations, nil
}
