package webserver

import "log"
import "net/http"
import "html/template"
import "trainstations_domain/stations"
import "trainstations_domain/storage/file"
import "trainstations_domain/storage/bartapi"

type PageData struct {
	SelectedStations []bartapi.TrainInfo
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
	stationRepository := file.NewStationStorage("stations.txt")
	allStations, _ := stationRepository.GetAll()

	abbr := r.URL.Query().Get("abbr")
	stationData := bartapi.TrainsFromBartAPI(abbr, "n")

	page := PageData{
		SelectedStations: stationData,
		AllStations:      allStations}

	tmpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Panic("Error occurred parsing template", err)
	}

	err = tmpl.Execute(rw, page)
	if err != nil {
		log.Panic("Error occurred writing template", err)
	}
}
