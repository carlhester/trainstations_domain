package webserver

import "fmt"
import "log"
import "net/http"
import "html/template"
import "trainstations_domain/stations"
import "trainstations_domain/lines"
import "trainstations_domain/storage/file"
import "trainstations_domain/storage/bartapi"

type PageData struct {
	SelectedStations []bartapi.TrainInfo
	AllStations      []stations.Station
	AllLines         []lines.Line
}

func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(rw http.ResponseWriter, r *http.Request) {
	stationRepository := file.NewStationStorage("./data/stations.txt")
	allStations, _ := stationRepository.GetAll()
	allLines, _ := lines.GetAllLines()

	abbr := r.URL.Query().Get("abbr")
	line := r.URL.Query().Get("line")

	fmt.Println(line)

	stationData := bartapi.TrainsFromBartAPI(abbr, "n")

	page := PageData{
		SelectedStations: stationData,
		AllStations:      allStations,
		AllLines:         allLines,
	}

	tmpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Panic("Error occurred parsing template", err)
	}

	err = tmpl.Execute(rw, page)
	if err != nil {
		log.Panic("Error occurred writing template", err)
	}
}
