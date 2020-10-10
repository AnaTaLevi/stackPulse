package main

import (
	"encoding/json"
	"fmt"
	"github.com/ip2location/ip2location-go"
	"log"
	"net/http"
	"os"
	"strconv"
)


var ipToCount = make(map[string]int)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	error := make(map[string]error)
	ip := r.URL.Query().Get("ip")
	if exceededLimitation(w, ip) {
		json.NewEncoder(w).Encode(http.StatusTooManyRequests)
		return
	}
	ipToCount[ip] += 1
	fileName, _ := os.LookupEnv("file")
	db, err := ip2location.OpenDB(fileName)

	if err != nil {
		error["error"] = err
		json.NewEncoder(w).Encode(error)
		return
	}
	results, err := db.Get_all(ip)

	if err != nil {
		error["error"] = err
		json.NewEncoder(w).Encode(error)
		return
	}

	location := make(map[string]string)
	location["country"] = results.Country_long
	location["city"] = results.City
	json.NewEncoder(w).Encode(location)
}

func exceededLimitation(ip string) bool {
	limitation, _ := os.LookupEnv("rate_limit")
	strLimitation, _ := strconv.Atoi(limitation)git
	if ipToCount[ip] >= strLimitation {
		return true
	}
	return false
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/v1/find-country", getLocation)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
