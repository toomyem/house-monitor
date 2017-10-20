package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const port = 9000

type temperatureReading struct {
	DeviceID  string `json:"device_id"`
	Timestamp int    `json:"ts"`
	Reading   string `json:"reading"`
}

func validateReading(data temperatureReading) error {
	if data.DeviceID == "" {
		return fmt.Errorf("'device_id' cannot be empty")
	}

	if data.Timestamp <= 0 {
		return fmt.Errorf("'ts' cannot be <= 0")
	}

	if data.Reading == "" {
		return fmt.Errorf("'reading' cannot be empty")
	}

	_, err := strconv.ParseFloat(data.Reading, 64)
	if err != nil {
		return fmt.Errorf("cannot parse 'reading' as float")
	}

	return nil
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Cannot read request body: %s", err)
		return
	}
	var data temperatureReading
	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Cannot parse body: %s", err)))
		return
	}
	err = validateReading(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid data: %s", err)))
		return
	}
	err = storeInDb(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Cannot store data in db: %s", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Received: %#v\n", data)
}

func main() {
	// err := initializeDb()
	// if err != nil {
	// 	log.Fatalf("Cannot connect to db: %s", err)
	// }

	http.HandleFunc("/temperature", temperatureHandler)

	log.Printf("App started on port %d\n", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil))
}
