package file

import "encoding/json"
import "os"
import "log"

//import "fmt"

import "trainstations_domain/stations"

type FileStationStorage struct {
	stations []stations.Station
}

// The constructor creates default
func NewFileStationStorage() *FileStationStorage {
	data := stations.Station{Abbr: "MONT", Name: "Montgomery"}
	storage := new(FileStationStorage)
	storage.Add(data)
	return nil
}

func (f *FileStationStorage) Add(station stations.Station) error {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, _ := json.Marshal(station)
	bytes := []byte(string(b))

	bytesWritten, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return nil
}
