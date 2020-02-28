package models

type Car struct {
	Vin         int `gorm:"primary_key"`
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

func (car Car) getVehicleType() string {
	return car.Vehicletype
}
func (truck Truck) getVehicleType() string {
	return truck.Vehicletype
}

func (car Car) getVehicleIdNumber() int {
	return car.Vin
}

func (truck Truck) getVehicleIdNumber() int {
	return truck.Vin
}
