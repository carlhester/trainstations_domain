package webserver

import "log"
import "net/http"
import "html/template"

import "trainstations_domain/stations"
import "trainstations_domain/storage/file"

type PageData struct {
	SelectedStations []stations.Station
	AllStations      []stations.Station
}

func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(rw http.ResponseWriter, r *http.Request) {
	stationRepository := file.NewStationStorage("test.txt")

	abbr := r.URL.Query().Get("abbr")

	allStationsData, _ := stationRepository.GetAll()
	stationData, _ := stationRepository.Get(abbr)

	page := PageData{
		SelectedStations: stationData,
		AllStations:      allStationsData}

	tmpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Panic("Error occurred parsing template", err)
	}

	err = tmpl.Execute(rw, page)
	if err != nil {
		log.Panic("Error occurred writing template", err)
	}
}
