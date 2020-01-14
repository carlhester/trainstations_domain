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
	stationRepository := file.NewStationStorage("test.txt")

	abbr := r.URL.Query().Get("abbr")

	bartdata := bartapi.TrainsFromBartAPI(abbr, "n")

	allStationsData, _ := stationRepository.GetAll()
	stationData := bartdata

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
