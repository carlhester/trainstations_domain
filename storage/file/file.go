package file

import "encoding/json"
import "os"
import "log"

import "trainstations_domain/stations"

type FileStationStorage struct {
	stations []stations.Station
	file     *os.File
}

func NewFileStationStorage(file string) *FileStationStorage {
	storage := new(FileStationStorage)
	newFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	storage.file = newFile
	return storage

}

func (f *FileStationStorage) Add(station stations.Station) error {
	b, _ := json.Marshal(station)
	bytes := []byte(string(b))
	bytesWritten, err := f.file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return nil
}
