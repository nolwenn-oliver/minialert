package server

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
	Name           string   `yaml:"name"`
	AlertMessage   string   `yaml:"alert_message"`
	AlertThreshold int      `yaml:"alert_threshold"`
	Notify         []string `yaml:"notify"`
}

func loadMetrics(path string) ([]Metric, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		log.Fatalf("Unable to decode metric config file %s\n", err)
	}

	return cfg.Metrics, nil
}

func getMetric(metrics []Metric, name string) (*Metric, error) {
	for _, m := range metrics {
		if m.Name == name {
			return &m, nil
		}
	}
	return nil, errors.New("Metric not found")
}
