package client

import (
	"bytes"
	"encoding/json"
	"log"
	"minialert/server"
	"net/http"
)

var serverURL = "http://localhost:"

func SendDataMetric(port string, name string, value string) {
	metricURL := serverURL + port + "/metric?name=" + name + "&value=" + value
	var requestBody bytes.Buffer
	res, err := http.Post(metricURL, "Content-Length: 0", &requestBody)
	if err != nil {
		log.Fatalf("Error making http request: %s\n", err)
	}
	log.Printf("Received status code: %d\n", res.StatusCode)
}

func GetAlertHistory(port string) {
	alertsURL := serverURL + port + "/alerts"
	res, err := http.Get(alertsURL)
	if err != nil {
		log.Fatalf("Error making http request: %s\n", err)
	}

	var alertsHistory []server.Alert
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&alertsHistory)
	if err != nil {
		log.Fatalf("Error while decoding server response: %s\n", err)
	}
	if len(alertsHistory) == 0 {
		log.Printf("Empty alerts history!")
	}
	for _, alert := range alertsHistory {
		log.Printf("[%s] %s: %s, notified: %s\n", alert.Timestamp, alert.AlertMessage, alert.Value, alert.Notify)
	}
}
