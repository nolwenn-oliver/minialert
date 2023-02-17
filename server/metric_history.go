package main

import (
	"fmt"
	"log"
	"time"
)

var AlertsHistory []Alert

type Alert struct {
	timestamp    string
	alertMessage string
	value        int
	//notify []string
}

func checkMetricAlert(name string, value int) {
	metricConfig, err := getMetric(name)
	if err != nil {
		log.Fatal("No configuration for metric ", name)
	}

	if value > metricConfig.AlertThreshold {
		alert := Alert{time.Now().Format(time.RFC850), metricConfig.AlertMessage, value}
		fmt.Printf("[%s][ALERT] %s: %d", alert.timestamp, alert.alertMessage, alert.value)
		AlertsHistory = append(AlertsHistory, alert)
	}
}
