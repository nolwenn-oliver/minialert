package server

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var AlertsHistory []Alert

type Alert struct {
	Timestamp    string
	AlertMessage string
	Value        string
	//notify []string
}

func checkMetricAlert(name string, value string) {
	metricConfig, err := getMetric(name)
	if err != nil {
		log.Fatal("No configuration for metric ", name)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("Invalid value for metric ", name)
	}

	log.Printf("[Metric]: %s -> %s\n", name, value)

	if intValue > metricConfig.AlertThreshold {
		alert := Alert{time.Now().Format(time.RFC850), metricConfig.AlertMessage, value}
		fmt.Printf("[%s][ALERT] %s: %s\n", alert.Timestamp, alert.AlertMessage, alert.Value)
		AlertsHistory = append(AlertsHistory, alert)
	}
}
