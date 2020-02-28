package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func init() {
	Db = gormsess()
	Db.Debug().AutoMigrate(&Car{})

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

func Daowelcome() {
	fmt.Println("this is dao function")
}

func InsertCar(car *Car) {

	db := Db.Create(car)

	fmt.Println("record inserted ")

	db.Close()

}
