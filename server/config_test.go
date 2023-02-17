package main_test

import (
	"reflect"
	"testing"
)

func TestLoadMetrics(t *testing.T) {
	metrics, _ := server.loadMetrics()
	want := []server.Metric{
		{"cpu", "High CPU usage", 80},
		{"battery", "Low battery", 80}}
	if !reflect.DeepEqual(metrics, want) {
		t.Errorf("got %q, wanted %q", metrics, want)
	}
}
