package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func sendDataMetric(w http.ResponseWriter, r *http.Request) {
	metric := r.URL.Query().Get("name")
	value := r.URL.Query().Get("value")
	if metric == "" {
		w.Header().Set("x-missing-field", "metric")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if value == "" {
		w.Header().Set("x-missing-field", "value")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	checkMetricAlert(metric, value)
	w.WriteHeader(http.StatusOK)
}
func getAlertHistory(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(AlertsHistory)
	if err != nil {
		fmt.Println(err)
	} else {
		io.WriteString(w, string(body))
	}
}

func main() {
	http.HandleFunc("/metric", sendDataMetric)
	http.HandleFunc("/alerts", getAlertHistory)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server closed\n")
	} else if err != nil {
		log.Fatal("Error starting server: %s\n", err)
	}
}
