package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)


func exceededLimitation(ip string) bool {
	limitation, _ := os.LookupEnv("rate_limit")
	strLimitation, _ := strconv.Atoi(limitation)
	if ipToCount[ip] >= strLimitation {
		return true
	}
	return false
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	error := make(map[string]error)
	ip := r.URL.Query().Get("ip")
	if exceededLimitation(ip) {
		json.NewEncoder(w).Encode(http.StatusTooManyRequests)
		return
	}
	ipToCount[ip] += 1
	fileName, _ := os.LookupEnv("file")
	err, results, done := getAddressFromDB(w, fileName, error, ip)
	if done {
		return
	}

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

