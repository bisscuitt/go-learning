package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// You can choose to only unmarshal some of the json data
// Create a data structure that only has fields for some of the data
type city struct {
	Latitude, Longitude float64
	City                string
}

type remoteError struct {
	Error   string
	Message string
	ErrorID int
}

type cities []city

func main() {
	var data cities
	//	rcvd := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	rcvd := `{"Error":"Error101","Message":"Failure in request string","ErrorID":12345}`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		var j_err remoteError
		_ = json.Unmarshal([]byte(rcvd), &j_err)
		log.Fatalln(err, j_err)
	}

	fmt.Println(data)
}
