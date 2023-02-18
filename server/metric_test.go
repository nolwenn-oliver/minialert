package server

import (
	"reflect"
	"testing"
)

func TestLoadMetrics(t *testing.T) {
	metrics, _ := loadMetrics()
	want := []Metric{
		{"cpu", "High CPU usage", 80},
		{"battery", "Low battery", 80}}
	if !reflect.DeepEqual(metrics, want) {
		t.Errorf("got %q, wanted %q", metrics, want)
	}
}
