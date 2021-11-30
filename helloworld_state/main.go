package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type stateData struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

var urlStr = "http://localhost:8089/v1.0/state/statestore"

func main() {
	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get(urlStr + "/mystate")
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		strVal := string(body)
		count := 0

		if strVal != "" {
			count, _ = strconv.Atoi(strVal)
			count++
		}

		stateObj := []stateData{stateData{Key: "mystate", Value: count}}
		stateData, _ := json.Marshal(stateObj)

		resp, _ = http.Post(urlStr, "application/json", bytes.NewBuffer(stateData))
		if count == 1 {
			fmt.Fprintf(w, "I've greeted you "+strconv.Itoa(count)+" time.")
		} else {
			fmt.Fprintf(w, "I've greeted you "+strconv.Itoa(count)+" times.")
		}
	})
	log.Fatal(http.ListenAndServe(":8088", nil))
}
