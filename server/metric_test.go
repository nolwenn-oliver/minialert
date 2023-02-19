package server

import (
	"reflect"
	"testing"
)

func TestLoadMetricsSuccess(t *testing.T) {
	want := []Metric{
		{"cpu", "High CPU usage", 80, []string{"morgan", "doran"}},
		{"battery", "Low battery", 80, []string{"nicolas"}}}
	metrics, _ := loadMetrics("../config/metric_config.yaml")
	if !reflect.DeepEqual(metrics, want) {
		t.Errorf("got %q, wanted %q", metrics, want)
	}
}

func TestGetMetricsSuccess(t *testing.T) {
	metric1 := Metric{"cpu", "High CPU usage", 80, []string{"morgan", "doran"}}
	metrics := []Metric{
		metric1,
		{"battery", "Low battery", 80, []string{"nicolas"}}}

	res, _ := getMetric(metrics, "cpu")
	if !reflect.DeepEqual(*res, metric1) {
		t.Errorf("got %q, wanted %q", res, metric1)
	}
}
