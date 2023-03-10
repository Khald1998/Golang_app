package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type State struct {
	ID  int    `json:"ID"`
	X   int    `json:"X"`
	Y   int    `json:"Y"`
	OP  string `json:"OP"`
	RES int    `json:"RES"`
}

var (
	username = "x"
	password = "x"
	endpoint = "x"
	port     = "x"
	db_name  = "x"
)

// var DNS string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, endpoint, port, db_name)
var DB *gorm.DB
var err error

const DNS = "username:password@tcp(maindb.czpld8fke1ht.us-east-1.rds.amazonaws.com:3306)/DB"

func InitRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/state/get", GetState).Methods("GET")
	router.HandleFunc("/state/get/{ID}", GetStateID).Methods("GET")
	router.HandleFunc("/state/add", CreateState).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
func initializeRouter() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&State{})
}
func initializeDB() {
	DB.

}
func main() {
	InitRouter()
	initializeRouter()
	initializeDB()
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func GetState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var state []State
	DB.Find(&state)
	json.NewEncoder(w).Encode(state)
}
func GetStateID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)
	var state State
	DB.First(&state, parameter["ID"])
	json.NewEncoder(w).Encode(state)
}
func CreateState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var state State
	json.NewDecoder(r.Body).Decode(&state)
	DB.Create(&state)
	json.NewEncoder(w).Encode(state)

}
