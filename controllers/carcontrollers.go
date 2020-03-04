package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../models"
)

func SetCar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("set method")
	m := make(map[string]string)
	car1 := &models.Car{}
	body, err := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &m)
	if err == nil {

		fmt.Println(m)
		car1.Vin, _ = strconv.Atoi(m["vin"])
		car1.Vehicletype = m["vehicletype"]
		car1.Make = m["make"]
		car1.Model = m["model"]
		car1.Year, _ = strconv.Atoi(m["year"])
		car1.Mileage, _ = strconv.Atoi(m["mileage"])

		models.InsertCar(car1)

	}
	fmt.Println(m)
}

func GetCars(res http.ResponseWriter, req *http.Request) {
	cars := models.Getcars()
	fmt.Println(cars)
}

func DeleteCar(res http.ResponseWriter, req *http.Request) {
	p1, _ := strconv.Atoi(req.URL.Query().Get("vin"))
	models.DeleteCar(p1)

}
func GetCarsbyVin(res http.ResponseWriter, req *http.Request) {
	vin, _ := strconv.Atoi(req.URL.Query().Get("vin"))
	fmt.Println(models.GetCarByVin(vin))
}
