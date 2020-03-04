package models

type Car struct {
	Vin         int
	Vehicletype string
	Make        string
	Model       string
	Year        int
	Mileage     int
}

type Cars struct {
	Vin         int
	Vehicletype string
	Make        string
	Model       string
	Year        int
	Mileage     int
}

type Truck struct {
	Vin         int
	Vehicletype string
	Make        string
	Model       string
	Year        int
	Load        int
}

type Vehicles interface {
	getVehicleType() string
	getVehicleIdNumber() int
}
