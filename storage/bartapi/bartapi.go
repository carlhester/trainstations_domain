package bartapi

import "log"
import "sort"
import "strconv"
import "net/url"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "trainstations_domain/trains"

func TrainsFromBartAPI(station string, dir string) []trains.TrainInfo {
	if station == "" || dir == "" {
		return nil
	}

	const APIKEY = "MW9S-E7SL-26DU-VV8V"
	params := url.Values{}
	params.Add("cmd", "etd")
	params.Add("orig", station)
	params.Add("key", APIKEY)
	params.Add("dir", dir)
	params.Add("json", "y")
	url := "http://api.bart.gov/api/etd.aspx?" + params.Encode()

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	rawdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	trainData := rawAPIDataIntoTrains(rawdata)
	var targetTrains []trains.TrainInfo

	for _, station := range trainData.Root.Station {
		for _, train := range station.Etd {
			for _, estimate := range train.Est {
				thisTrain := trains.TrainInfo{
					Dest:    train.Destination,
					Minutes: convertMinutesToInt(estimate.Minutes),
					Line:    estimate.Color,
					Points:  0,
				}
				targetTrains = append(targetTrains, thisTrain)
			}
		}
	}
	sortedTargetTrains := sortTrainsByMinutes(targetTrains)
	return sortedTargetTrains
}

func rawAPIDataIntoTrains(raw []byte) *RawTrainData {
	var trainData RawTrainData
	_ = json.Unmarshal([]byte(raw), &trainData)
	return &trainData
}

func sortTrainsByMinutes(targets []trains.TrainInfo) []trains.TrainInfo {
	sort.Slice(targets, func(i, j int) bool { return targets[i].Minutes < targets[j].Minutes })
	return targets
}

func convertMinutesToInt(minutes string) int {
	if minutes == "Leaving" {
		minutes = "0"
	}
	i, err := strconv.Atoi(minutes)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
