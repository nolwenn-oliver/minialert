package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func sendDataMetric(w http.ResponseWriter, r *http.Request) {
	checkMetricAlert("cpu", 90)
	//io.WriteString(w, AlertsHistory[0].alertMessage)
}
func getAlertHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/", sendDataMetric)
	http.HandleFunc("/hello", getAlertHistory)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server closed\n")
	} else if err != nil {
		log.Fatal("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
