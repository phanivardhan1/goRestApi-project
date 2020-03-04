package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func init() {
	Db = gormsess()
	Db.Debug().AutoMigrate(&Car{}, &Truck{})

}

func gormsess() *gorm.DB {
	Db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=pgadmin sslmode=disable")
	fmt.Println("gorm session is created")
	if err != nil {
		fmt.Println(err.Error())
	}
	return Db

}

// func sess() *sql.DB {
// 	fmt.Println("init in dao is executed")
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//defer db.Close()
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Successfully connected!")

// 	return db
// }

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pgadmin"
	dbname   = "postgres"
)

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

func Daowelcome() {
	fmt.Println("this is dao function")
}

func InsertCar(car *Car) {
	Db.Create(car)
	fmt.Println("record inserted ")

}

func Getcars() []Car {
	var Cars []Car
	Db.Find(&Cars)
	return Cars

}
func DeleteCar(vin int) {

	Db.Where("vin = ?", vin).Delete(Car{})
	fmt.Println("the car is deleted")
}

func GetCarByVin(vin int) Cars {
	var car Cars
	Db.Where("vin = ?", vin).Find(&car)
	fmt.Println("get carby Id", vin)
	return car
}

func InsertTruck(truck *Truck) {
	Db.Create(truck)
	fmt.Println("record inserted ")
}

func GetTrucksFromDb() []Truck {
	var truck []Truck
	Db.Find(&truck)

	fmt.Println(truck)
	return truck

}

func GetTucksByVinFromDb(vin int) Truck {
	var truck Truck
	Db.Where("vin = ?", vin).Find(&truck)

	fmt.Println(truck)
	return truck
}
