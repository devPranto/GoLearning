package main

import (
	"api-project/package/model"
	// "encoding/json"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
	"time"
	// "github.com/gorilla/mux"
)

func main() {
	
	address := "http://api.aladhan.com/v1/calendar?latitude=23.777176&longitude=90.399452&method=2&month=1&year=2023"
	var response model.ResponseStruct
	getJson(address, &response)
	for _,value := range response.Data{
		fmt.Printf("%v %v\n",value.Timings,value.Date.Readable)
	}
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
