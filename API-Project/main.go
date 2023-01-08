package main

import (
	"api-project/package/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var Response model.ResponseStruct

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getTiming).Methods("GET")
	//router.HandleFunc("/{date}", getTiming).Methods("GET")
	//router.HandleFunc("/", getTiming).Methods("GET")
	//router.HandleFunc("/", getTiming).Methods("GET")
	fetchAPI()
	log.Fatal(http.ListenAndServe(":9010", router))

}

func fetchAPI() {
	address := "http://api.aladhan.com/v1/calendar?latitude=23.777176&longitude=90.399452&method=2&month=1&year=2023"

	getJson(address, &Response)

	dataFrame := &model.PrayerTiming{}
	for index, value := range Response.Data {
		dataFrame.ID = index + 1
		dataFrame.Fajr = value.Timings.Fajr
		dataFrame.Asr = value.Timings.Asr
		dataFrame.Dhuhr = value.Timings.Dhuhr
		dataFrame.Maghrib = value.Timings.Maghrib
		dataFrame.Isha = value.Timings.Isha
		dataFrame.Date = value.Date.Readable
		dataFrame.Creat()
	}
	//fmt.Printf("%+v \n", dataFrame)
}

func getJson(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
func getTiming(writer http.ResponseWriter, request *http.Request) {
	t := time.Now()
	r := model.FindByDate(int(t.Day()))
	re, _ := json.Marshal(r)
	writer.Write(re)
}
