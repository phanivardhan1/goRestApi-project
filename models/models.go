package models

type Car struct {
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

type User struct {
	Userid     int    `gorm:"primary_key;AUTO_INCREMENT"`
	FirstName  string `gorm:"type:varchar(100)"`
	LastName   string `gorm:"type:varchar(100)"`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Address    `gorm:"association_foreignkey:AddressId"`
	Preference `gorm:"association_foreignkey:PrefernceId"`
}

type Address struct {
	AddressId int    `gorm:"primary_key;AUTO_INCREMENT"`
	Street    string `gorm:"type:varchar(100)"`
	City      string `gorm:"type:varchar(100)"`
	State     string `gorm:"type:varchar(100)"`
	Zipcode   string `gorm:"type:varchar(100)"`
}
type Preference struct {
	PrefernceId int    `gorm:"primary_key;AUTO_INCREMENT"`
	Make        string `gorm:"type:varchar(100)"`
	Startyear   int
	EndYear     int
}

type Vehicles interface {
	getVehicleType() string
	getVehicleIdNumber() int
}
