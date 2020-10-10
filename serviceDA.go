package main

import (
	"encoding/json"
	"github.com/ip2location/ip2location-go"
	"net/http"
)


func getAddressFromDB(w http.ResponseWriter, fileName string, error map[string]error, ip string) (error, ip2location.IP2Locationrecord, bool) {
	db, err := ip2location.OpenDB(fileName)

	if err != nil {
		error["error"] = err
		json.NewEncoder(w).Encode(error)
		return nil, ip2location.IP2Locationrecord{}, true
	}
	results, err := db.Get_all(ip)
	return err, results, false
}

