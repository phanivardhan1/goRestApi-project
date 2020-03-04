package main

import (
	"fmt"
	"net/http"

	"./controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//###############################################################################################
func main() {
	fmt.Println("Hi this is a rest Application")

	//dao.Insertcar()

	mux := mux.NewRouter()
	mux.HandleFunc("/home", welcome)
	mux.HandleFunc("/getAllvehicles", controllers.GetAllVehicles)
	mux.HandleFunc("/getvehiclesofmake", controllers.GetVehiclesofMake)
	mux.HandleFunc("/getallvehiclesofmake", controllers.GetAllVehiclesofMake)
	mux.HandleFunc("/getallvehiclesinrange", controllers.GetAllVehiclesinRange)
	mux.HandleFunc("/setCars", controllers.SetCar)
	mux.HandleFunc("/getCars", controllers.GetCars)
	mux.HandleFunc("/getCarByVin", controllers.GetCarsbyVin)
	mux.HandleFunc("/deleteCar", controllers.DeleteCar)
	mux.HandleFunc("/setTrucks", controllers.InserTruck).Methods("POST")
	mux.HandleFunc("/getTrucks", controllers.GetTrucks).Methods("GET")
	mux.HandleFunc("/getTrucksByVin", controllers.GetTrucksByVin).Methods("GET")
	http.ListenAndServe(":8080", mux)
}

// #############################################################################################
func welcome(res http.ResponseWriter, req *http.Request) {
	fmt.Println("welcome function")

}
