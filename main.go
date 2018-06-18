package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Car struct {
	ID         string      `json:"id,omitempty"`
	Carname    string      `json:"carname,omitempty"`
	Cardesc    string      `json:"cardesc,omitempty"`
	EngineType *EngineType `json:"enginetype,omitempty"`
}

type EngineType struct {
	EngineNumber string `json:"enginenumber,omitempty"`
	FuelType     string `json:"fueltype,omitempty"`
}

var cars []Car

func GetCarEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}

func GetCarsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(cars)
}

func CreateCarEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var car Car
	_ = json.NewDecoder(req.Body).Decode(&car)
	car.ID = params["id"]
	cars = append(cars, car)
	json.NewEncoder(w).Encode(cars)
}

func DeleteCarEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(cars)
}

func main() {
	router := mux.NewRouter()
	cars = append(cars, Car{ID: "1", Carname: "Gallardo", Cardesc: "Fast and awesome", EngineType: &EngineType{EngineNumber: "AQWS12354", FuelType: "Petrol"}})
	cars = append(cars, Car{ID: "2", Carname: "360 Spider", Cardesc: "it is also good"})
	router.HandleFunc("/cars", GetCarsEndpoint).Methods("GET")
	router.HandleFunc("/cars/{id}", GetCarEndpoint).Methods("GET")
	router.HandleFunc("/cars/{id}", CreateCarEndpoint).Methods("POST")
	router.HandleFunc("/cars/{id}", DeleteCarEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}
