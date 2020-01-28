package webserver

//import "fmt"
import "log"
import "net/http"
import "html/template"
import "trainstations_domain/stations"
import "trainstations_domain/lines"
import "trainstations_domain/trains"
import "trainstations_domain/scoring"
import "trainstations_domain/storage/file"
import "trainstations_domain/storage/bartapi"

type PageData struct {
	SelectedStations []trains.TrainInfo
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

	lineRepository := file.NewLineStorage("./data/lines.txt")
	allLines, _ := lineRepository.GetAll()

	_ = r.ParseForm()
	abbr := r.URL.Query().Get("abbr")
	lines := r.Form["line"]

	stationData := bartapi.TrainsFromBartAPI(abbr, "n")

	filteredTrains := filterDestinationByLine(stationData, lines)
	scoredTrains := scoring.Score(filteredTrains)

	page := PageData{
		//SelectedStations: filteredTrains,
		SelectedStations: scoredTrains,
		AllStations:      allStations,
		AllLines:         allLines,
	}

	templates := []string{"./ui/html/index.html", "./ui/html/templates/base.tmpl"}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Panic("Error occurred parsing template", err)
	}

	err = tmpl.Execute(rw, page)
	if err != nil {
		log.Panic("Error occurred writing template", err)
	}
}

func filterDestinationByLine(trainsToFilter []trains.TrainInfo, lines []string) []trains.TrainInfo {
	var trainsMatchLine []trains.TrainInfo
	for _, train := range trainsToFilter {
		for _, line := range lines {
			if train.Line == line {
				trainsMatchLine = append(trainsMatchLine, train)
			}
		}
	}
	return trainsMatchLine
}
