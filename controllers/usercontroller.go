package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"../utilities"

	"../models"
)

func InsertUsers(res http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	m := make(map[string]interface{})
	Body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		json.Unmarshal([]byte(Body), &m)
		fmt.Println("the map is", m)
		user.FirstName = m["FirstName"].(string)
		user.LastName = m["LastName"].(string)
		user.Email = m["Email"].(string)
		fmt.Println(reflect.TypeOf(m["Address"]))
		fmt.Println(reflect.TypeOf(m["Preference"]))

		Address := m["Address"].(map[string]interface{})
		user.Street = Address["Street"].(string)
		user.City = Address["City"].(string)
		user.State = Address["state"].(string)
		user.Zipcode = Address["Zipcode"].(string)

		Preference := m["Preference"].(map[string]interface{})
		user.Make = Preference["Make"].(string)

		fmt.Println(reflect.TypeOf(Preference["Startyear"]))
		user.Startyear, _ = strconv.Atoi(Preference["Startyear"].(string))
		user.EndYear, _ = strconv.Atoi(Preference["EndYear"].(string))

		fmt.Println("The user struct is", user)
	}

	models.InserUsersInDb(user)
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	users := models.GetUsersFromDb()
	utilities.RespondJSON(res, 200, users)
}

func GetCarsByUserPreference(res http.ResponseWriter, req *http.Request) {

	firstname := req.URL.Query().Get("FirstName")
	lastname := req.URL.Query().Get("LastName")
	email := req.URL.Query().Get("Email")

	fmt.Println(firstname, lastname, email)

	models.GetCarsFromDbByUserPreference(firstname, lastname, email)

}
