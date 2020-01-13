package webserver

import "log"
import "net/http"
import "html/template"

import "trainstations_domain/stations"
import "trainstations_domain/storage/file"

type PageData struct {
	Stations []stations.Station
}

func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Panic("Error occurred parsing template", err)
	}

	storageFile := "test.txt"
	stationRepository := file.NewStationStorage(storageFile)
	allStations, _ := stationRepository.GetAll()

	page := PageData{
		Stations: allStations,
	}

	err = tmpl.Execute(rw, page)
	if err != nil {
		log.Panic("Error occurred writing template", err)
	}
}
