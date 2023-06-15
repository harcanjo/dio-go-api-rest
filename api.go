package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ParkingSpot struct {
	ID       string `json:"id"`
	Number   int    `json:"number"`
	Occupied bool   `json:"occupied"`
}

type Car struct {
	ID     string `json:"id"`
	Plate  string `json:"plate"`
	Color  string `json:"color"`
	SpotID string `json:"spotId"`
}

var parkingSpots []ParkingSpot
var cars []Car

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/parkingSpots", getAvailableParkingSpots).Methods("GET")
	router.HandleFunc("/parkingSpots", createParkingSpot).Methods("POST")
	router.HandleFunc("/cars", parkCar).Methods("POST")
	router.HandleFunc("/cars/{id}", retrieveCar).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAvailableParkingSpots(w http.ResponseWriter, r *http.Request) {
	availableSpots := make([]ParkingSpot, 0)
	for _, spot := range parkingSpots {
		if !spot.Occupied {
			availableSpots = append(availableSpots, spot)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(availableSpots)
}

func createParkingSpot(w http.ResponseWriter, r *http.Request) {
	var newSpot ParkingSpot
	err := json.NewDecoder(r.Body).Decode(&newSpot)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newSpot.ID = fmt.Sprintf("spot%d", len(parkingSpots)+1)
	newSpot.Occupied = false

	parkingSpots = append(parkingSpots, newSpot)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSpot)
}

func parkCar(w http.ResponseWriter, r *http.Request) {
	var newCar Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Find an available parking spot
	var spot *ParkingSpot
	for i := range parkingSpots {
		if !parkingSpots[i].Occupied {
			spot = &parkingSpots[i]
			break
		}
	}

	if spot == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	newCar.ID = fmt.Sprintf("car%d", len(cars)+1)
	newCar.SpotID = spot.ID

	spot.Occupied = true
	cars = append(cars, newCar)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

func retrieveCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carID := params["id"]

	for i, car := range cars {
		if car.ID == carID {
			// Free up the parking spot
			for j := range parkingSpots {
				if parkingSpots[j].ID == car.SpotID {
					parkingSpots[j].Occupied = false
					break
				}
			}

			// Remove the car from the list
			cars = append(cars[:i], cars[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
