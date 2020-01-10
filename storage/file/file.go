package file

import "encoding/json"
import "os"
import "log"

import "trainstations_domain/stations"

type FileStationStorage struct {
	stations []stations.Station
}

func NewFileStationStorage(file string) (*FileStationStorage, *os.File) {
	storage := new(FileStationStorage)

	newFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	return storage, newFile
}

func (f *FileStationStorage) Add(fp *os.File, station stations.Station) error {
	b, _ := json.Marshal(station)
	bytes := []byte(string(b))
	bytesWritten, err := fp.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return nil
}
