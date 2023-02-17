package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Metrics []Metric `yaml:"metrics"`
}

type Metric struct {
	Name           string `yaml:"name"`
	AlertMessage   string `yaml:"alert_message"`
	AlertThreshold int    `yaml:"alert_threshold"`
}

func loadMetrics() ([]Metric, error) {
	f, err := os.Open("config/metric_config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cfg.Metrics, nil
}

func getMetric(name string) (*Metric, error) {
	metrics, _ := loadMetrics()
	for _, m := range metrics {
		if m.Name == name {
			return &m, nil
		}
	}
	return nil, errors.New("No configuration for metric")
}
