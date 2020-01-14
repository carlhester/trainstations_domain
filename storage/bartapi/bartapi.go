package bartapi

import "log"
import "net/url"
import "net/http"
import "io/ioutil"
import "encoding/json"

func TrainsFromBartAPI(station string, dir string) []TrainInfo {
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
	var targetTrains []TrainInfo

	for _, station := range trainData.Root.Station {
		for _, train := range station.Etd {
			for _, estimate := range train.Est {
				thisTrain := TrainInfo{
					Dest:    train.Destination,
					Minutes: estimate.Minutes,
					Line:    estimate.Color,
				}
				targetTrains = append(targetTrains, thisTrain)
			}
		}
	}
	return targetTrains
}

func rawAPIDataIntoTrains(raw []byte) *RawTrainData {
	var trainData RawTrainData
	_ = json.Unmarshal([]byte(raw), &trainData)
	return &trainData
}

type TrainInfo struct {
	Dest    string
	Minutes string
	Line    string
}

type Estimates []struct {
	Minutes     string `json:"minutes"`
	Direction   string `json:"direction"`
	Length      int    `json:"length"`
	Color       string `json:"color"`
	Hexcolor    string `json:"hexcolor"`
	Bikeflag    int    `json:"bikeflag"`
	Delay       int    `json:"delay"`
	Carflag     int    `json:"carflag"`
	Cancelflag  int    `json:"cancelflag"`
	Dynamicflag int    `json:"dynamicflag"`
}

type Etd []struct {
	Destination  string    `json:"destination"`
	Abbreviation string    `json:"abbreviation"`
	Limited      int       `json:"limited"`
	Est          Estimates `json:"estimate"`
}

type Station []struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
	Etd  Etd    `json:"etd"`
}

type Uri struct {
	Cdata string `json:"#cdata-section"`
}

type Root struct {
	Id      int     `json:"@id"`
	Uri     Uri     `json:"uri"`
	Date    string  `json:"date"`
	Time    string  `json:"time"`
	Station Station `json:"station"`
	Message string  `json:"message"`
}

type Xml struct {
	Version  string `json:"@version"`
	Encoding string `json:"@encoding"`
}

type RawTrainData struct {
	Xml  Xml  `json:"?xml"`
	Root Root `json:"root"`
}
