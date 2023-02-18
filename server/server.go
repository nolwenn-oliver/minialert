package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var DefaultPort = "8080"

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

	go checkMetricAlert(metric, value)
	w.WriteHeader(http.StatusCreated)
}
func getAlertHistory(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(AlertsHistory)
	if err != nil {
		fmt.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(body))
	}
}

func LaunchServer(port string) {
	http.HandleFunc("/metric", sendDataMetric)
	http.HandleFunc("/alerts", getAlertHistory)
	fmt.Println("[minialert server started]")

	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server closed\n")
	} else if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
