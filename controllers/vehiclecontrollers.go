package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"../models"
)

func CarsList(c chan []models.Car) {
	var carslist []models.Car
	c1 := models.Car{
		Vin:         1003,
		Vehicletype: "sedan",
		Make:        "mazda",
		Model:       "cx3",
		Year:        2019,
		Mileage:     27,
	}
	c2 := models.Car{
		Vin:         1001,
		Vehicletype: "sedan",
		Make:        "Honda",
		Model:       "civic",
		Year:        2016,
		Mileage:     32,
	}
	c3 := models.Car{
		Vin:         1002,
		Vehicletype: "suv",
		Make:        "toyota",
		Model:       "rav4",
		Year:        2018,
		Mileage:     26,
	}
	c4 := models.Car{
		Vin:         1003,
		Vehicletype: "sedan",
		Make:        "mazda",
		Model:       "cx3",
		Year:        2019,
		Mileage:     27,
	}
	c5 := models.Car{
		Vin:         1004,
		Vehicletype: "suv",
		Make:        "honda",
		Model:       "pilot",
		Year:        2020,
		Mileage:     25,
	}
	carslist = append(carslist, c1, c2, c3, c4, c5)
	c <- carslist
}

func TruckList(c chan []models.Truck) {
	var trucklist []models.Truck
	t1 := models.Truck{
		Vin:         2000,
		Vehicletype: "10wheel",
		Make:        "toyota",
		Model:       "camrytruck",
		Year:        2016,
		Load:        18,
	}
	t2 := models.Truck{
		Vin:         2001,
		Vehicletype: "10wheel",
		Make:        "Honda",
		Model:       "civictruck",
		Year:        2017,
		Load:        32,
	}
	t3 := models.Truck{
		Vin:         2002,
		Vehicletype: "20wheel",
		Make:        "toyota",
		Model:       "rav4truck",
		Year:        2018,
		Load:        17,
	}
	t4 := models.Truck{
		Vin:         2003,
		Vehicletype: "20wheel",
		Make:        "mazda",
		Model:       "cx3truck",
		Year:        2019,
		Load:        21,
	}
	t5 := models.Truck{
		Vin:         2004,
		Vehicletype: "20wheel",
		Make:        "honda",
		Model:       "pilottruck",
		Year:        2020,
		Load:        21,
	}

	trucklist = append(trucklist, t1, t2, t3, t4, t5)
	c <- trucklist
}

func GetVehicles() map[string]interface{} {

	c := make(chan []models.Car)
	d := make(chan []models.Truck)
	m := make(map[string]interface{})
	go CarsList(c)
	go TruckList(d)
	carlist, trucklist := <-c, <-d
	m["car"] = carlist
	m["truck"] = trucklist

	return m
}
func GetAllVehicles(res http.ResponseWriter, req *http.Request) {

	r := req.URL.Query().Get("vehicle")
	m := GetVehicles()
	if r == "car" {
		fmt.Println(m["car"])
	} else if r == "truck" {
		fmt.Println(m["truck"])
	}
}

func GetVehiclesofMake(res http.ResponseWriter, req *http.Request) {

	p1 := req.URL.Query().Get("vehicle")
	p2 := req.URL.Query().Get("make")

	//var vehiclesofmake []car
	fmt.Println(p1, p2)

	m := GetVehicles()
	//fmt.Println(m)
	if p1 == "car" {
		vehiclelist := m[p1]
		for _, v := range vehiclelist.([]models.Car) {
			if v.Make == p2 {
				fmt.Println(v)
			}

		}
	} else if p1 == "truck" {
		vehiclelist := m[p1]
		for _, v := range vehiclelist.([]models.Truck) {
			if v.Make == p2 {
				fmt.Println(v)
			}

		}

	}

}

func GetAllVehiclesofMake(res http.ResponseWriter, req *http.Request) {

	p1 := req.URL.Query().Get("make")
	fmt.Println("param 1 is", p1)
	m := GetVehicles()
	fmt.Println(m)
	var vehicleofmake []models.Vehicles

	vehicleslist, ok := m["car"]
	if ok == true {
		for _, v := range vehicleslist.([]models.Car) {
			if v.Make == p1 {
				vehicleofmake = append(vehicleofmake, v)
			}
		}
	}
	vehiclelist, ok := m["truck"]
	if ok == true {
		for _, v := range vehiclelist.([]models.Truck) {
			if v.Make == p1 {
				vehicleofmake = append(vehicleofmake, v)
			}
		}
	}

}

func GetAllVehiclesinRange(res http.ResponseWriter, req *http.Request) {

	p1, _ := strconv.Atoi(req.URL.Query().Get("start"))
	p2, _ := strconv.Atoi(req.URL.Query().Get("end"))

	m := GetVehicles()
	fmt.Println(m)
	var vehicleofmake []models.Vehicles

	vehicleslist, ok := m["car"]
	if ok == true {
		for _, v := range vehicleslist.([]models.Car) {
			if v.Year <= p2 && v.Year >= p1 {
				vehicleofmake = append(vehicleofmake, v)
			}
		}
	}
	vehiclelist, ok := m["truck"]
	if ok == true {
		for _, v := range vehiclelist.([]models.Truck) {
			if v.Year <= p2 && v.Year >= p1 {
				vehicleofmake = append(vehicleofmake, v)
			}
		}
	}

	fmt.Println(vehicleofmake)
}
