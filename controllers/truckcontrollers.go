package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../utilities"

	"../models"
)

func InserTruck(res http.ResponseWriter, req *http.Request) {
	truck := &models.Truck{}
	var m map[string]string
	Body, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(Body, &m)
	if err == nil {
		fmt.Println(m)
		truck.Vin, _ = strconv.Atoi(m["Vin"])
		truck.Vehicletype = m["Vehicletype"]
		truck.Make = m["Make"]
		truck.Model = m["Model"]
		truck.Year, _ = strconv.Atoi(m["Year"])
		truck.Load, _ = strconv.Atoi(m["Load"])

	} else {
		fmt.Println("error in the body")
	}

	models.InsertTruck(truck)

}

func GetTrucks(res http.ResponseWriter, req *http.Request) {
	trucks := models.GetTrucksFromDb()
	utilities.RespondJSON(res, 200, trucks)
}

func GetTrucksByVin(res http.ResponseWriter, req *http.Request) {
	vin, _ := strconv.Atoi(req.URL.Query().Get("vin"))

	fmt.Println("vin is ", vin)
	truck := models.GetTucksByVinFromDb(vin)

	utilities.RespondJSON(res, 200, truck)

}
func UpdateTrucks(res http.ResponseWriter, req *http.Request) {
	//truck := models.UpdateTruckInDb()
	//utilities.RespondJSON(res, 200, truck)
}
