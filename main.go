package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const port = 9000

type temperature struct {
	Reading float64
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
	var data temperature
	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot parse body"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Received: %#v\n", data)
}

func main() {
	http.HandleFunc("/temperature", temperatureHandler)

	log.Printf("App started on port %d\n", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil))
}
