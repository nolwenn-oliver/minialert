package server

import (
	"log"
	"strconv"
	"time"
)

var AlertsHistory []Alert
var configMetricPath = "config/metric_config.yaml"

type Alert struct {
	Timestamp    string
	AlertMessage string
	Value        string
	Notify       []string
}

func checkMetricAlert(name string, value string) {
	metrics, err := loadMetrics(configMetricPath)
	if err != nil {
		log.Fatalf("Impossible to open metric config file %s\n", name)
	}
	metricConfig, err := getMetric(metrics, name)
	if err != nil {
		log.Fatalf("No configuration for metric %s\n", name)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Invalid value for metric %s\n", name)
	}

	log.Printf("[Metric]: %s -> %s\n", name, value)

	if intValue > metricConfig.AlertThreshold {
		alert := Alert{time.Now().Format(time.RFC850), metricConfig.AlertMessage, value, metricConfig.Notify}
		log.Printf("[%s][ALERT] %s: %s, notifying: %s\n", alert.Timestamp, alert.AlertMessage, alert.Value, alert.Notify)
		AlertsHistory = append(AlertsHistory, alert)
	}
}
