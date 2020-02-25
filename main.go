package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type vehicles interface {
	getVehicleType() string
	getVehicleIdNumber() int
}

type car struct {
	vin         int
	vehicletype string
	make        string
	model       string
	year        int
	mileage     int
}

type truck struct {
	vin         int
	vehicletype string
	make        string
	model       string
	year        int
	load        int
}

func (car car) getVehicleType() string {
	return car.vehicletype
}
func (truck truck) getVehicleType() string {
	return truck.vehicletype
}

func (car car) getVehicleIdNumber() int {
	return car.vin
}

func (truck truck) getVehicleIdNumber() int {
	return truck.vin
}

func carslist(c chan []car) {
	var carslist []car
	c1 := car{
		vin:         1000,
		vehicletype: "sedan",
		make:        "toyota",
		model:       "camry",
		year:        2016,
		mileage:     26,
	}
	c2 := car{
		vin:         1001,
		vehicletype: "sedan",
		make:        "Honda",
		model:       "civic",
		year:        2016,
		mileage:     32,
	}
	c3 := car{
		vin:         1002,
		vehicletype: "suv",
		make:        "toyota",
		model:       "rav4",
		year:        2018,
		mileage:     26,
	}
	c4 := car{
		vin:         1003,
		vehicletype: "sedan",
		make:        "mazda",
		model:       "cx3",
		year:        2019,
		mileage:     27,
	}
	c5 := car{
		vin:         1004,
		vehicletype: "suv",
		make:        "honda",
		model:       "pilot",
		year:        2020,
		mileage:     25,
	}
	carslist = append(carslist, c1, c2, c3, c4, c5)
	c <- carslist
}

func trucklist(c chan []truck) {
	var trucklist []truck
	t1 := truck{
		vin:         2000,
		vehicletype: "10wheel",
		make:        "toyota",
		model:       "camrytruck",
		year:        2016,
		load:        18,
	}
	t2 := truck{
		vin:         2001,
		vehicletype: "10wheel",
		make:        "Honda",
		model:       "civictruck",
		year:        2017,
		load:        32,
	}
	t3 := truck{
		vin:         2002,
		vehicletype: "20wheel",
		make:        "toyota",
		model:       "rav4truck",
		year:        2018,
		load:        17,
	}
	t4 := truck{
		vin:         2003,
		vehicletype: "20wheel",
		make:        "mazda",
		model:       "cx3truck",
		year:        2019,
		load:        21,
	}
	t5 := truck{
		vin:         2004,
		vehicletype: "20wheel",
		make:        "honda",
		model:       "pilottruck",
		year:        2020,
		load:        21,
	}
	trucklist = append(trucklist, t1, t2, t3, t4, t5)
	c <- trucklist
}

func main() {
	fmt.Println("Hi this is a rest Application")
	mux := mux.NewRouter()
	mux.HandleFunc("/home", welcome)
	mux.HandleFunc("/getAllvehicles", getAllvehicles)
	mux.HandleFunc("/getvehiclesofmake", getvehiclesofmake)
	mux.HandleFunc("/getallvehiclesofmake", getAllVehiclesofMake)

	http.ListenAndServe(":8080", mux)
}
func welcome(res http.ResponseWriter, req *http.Request) {
	fmt.Println("welcome function")

}

func getvehicles() map[string]interface{} {

	c := make(chan []car)
	d := make(chan []truck)
	m := make(map[string]interface{})
	go carslist(c)
	go trucklist(d)
	carlist, trucklist := <-c, <-d
	m["car"] = carlist
	m["truck"] = trucklist

	return m
}
func getAllvehicles(res http.ResponseWriter, req *http.Request) {

	r := req.URL.Query().Get("vehicle")
	m := getvehicles()
	if r == "car" {
		fmt.Println(m["car"])
	} else if r == "truck" {
		fmt.Println(m["truck"])
	}
}

func getvehiclesofmake(res http.ResponseWriter, req *http.Request) {

	p1 := req.URL.Query().Get("vehicle")
	p2 := req.URL.Query().Get("make")

	//var vehiclesofmake []car
	fmt.Println(p1, p2)

	m := getvehicles()
	//fmt.Println(m)
	if p1 == "car" {
		vehiclelist := m[p1]
		for _, v := range vehiclelist.([]car) {
			if v.make == p2 {
				fmt.Println(v)
			}

		}
	} else if p1 == "truck" {
		vehiclelist := m[p1]
		for _, v := range vehiclelist.([]truck) {
			if v.make == p2 {
				fmt.Println(v)
			}

		}

	}

}

func getAllVehiclesofMake(res http.ResponseWriter, req *http.Request) {

	p1 := req.URL.Query().Get("make")
	fmt.Println("param 1 is", p1)
	m := getvehicles()
	fmt.Println(m)
	var vehicleofmake []vehicles

	vehicleslist, ok := m["car"]
	if ok == true {
		for _, v := range vehicleslist.([]car) {
			if v.make == p1 {
				vehicleofmake = append(vehicleofmake, v)
			}
		}
	}
	vehiclelist, ok := m["truck"]
	if ok == true {
		for _, v := range vehiclelist.([]truck) {
			if v.make == p1 {
				vehicleofmake = append(vehicleofmake, v)
			}
		}
	}

}
