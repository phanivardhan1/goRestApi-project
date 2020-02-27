package dao

import (
	"database/sql"
	"fmt"
)

func sess() *sql.DB {
	fmt.Println("init in dao is executed")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return db
}

type car struct {
	vin         int
	vehicletype string
	make        string
	model       string
	year        int
	mileage     int
}

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

func Getcarslist() []car {

	var result []car
	cars := car{}
	qry := "SELECT * FROM public.car"
	db := sess()
	rows, _ := db.Query(qry)
	columns, _ := rows.Columns()

	fmt.Println(columns)

	for rows.Next() {
		_ = rows.Scan(cars.vin, cars.vehicletype, cars.make, cars.model, cars.year, cars.mileage)
		result = append(result, cars)
	}

	return result
}
