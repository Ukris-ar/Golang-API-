package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type CityStruct struct {
	ID				string
	Name			string
	CountryCode		string
	District		string
	Population		string
}

// Init City struct
var city []CityStruct

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "World"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	
    if err != nil {
        panic(err.Error())
	}
	
    return db
}

func getCountryAll(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	result, err := db.Query("SELECT * FROM city;")

	if err != nil {
        panic(err.Error())
	}
	
	fetchData := CityStruct{}
	for result.Next() {
        err = result.Scan(&fetchData.ID,&fetchData.Name,&fetchData.CountryCode,&fetchData.District,&fetchData.Population)
        if err != nil {
            panic(err.Error())
		}

		// Assign to Struct
		city = append(city, fetchData);
	}

	json.NewEncoder(w).Encode(city)
 }

func homePage(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprint(w, "Welcome to the HomePage!")
}
func handleRequest() {
	http.HandleFunc("/", homePage) 
	http.HandleFunc("/getCountry", getCountryAll) 
    http.ListenAndServe(":8080", nil) 
}

func main() {
	handleRequest()
}