package main

import (
	"fmt"
	"log"
	"net/http"
)


var ipToCount = make(map[string]int)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/v1/find-country", getLocation)
	log.Fatal(http.ListenAndServe(":10000", nil))
}


func main() {
	handleRequests()
}
