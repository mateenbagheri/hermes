// Package db provides configuration for connecting to databases.
package db

import (
	"os"
)

// InfluxConfig is an interface for getting the configuration details required to connect to InfluxDB.
type InfluxConfig interface {
	Address() string
	AuthToken() string
	Organization() string
	Bucket() string
}

// influxConfig implements InfluxConfig and holds the configuration details for InfluxDB.
type influxConfig struct {
	address      string
	authToken    string
	organization string
	bucket       string
}

// NewInfluxConfig returns a new instance of influxConfig populated with environment variable values.
func NewInfluxConfig() InfluxConfig {
	return &influxConfig{
		address:      os.Getenv("INFLUX_ADDRESS"),
		authToken:    os.Getenv("INFLUX_TOKEN"),
		organization: os.Getenv("INFLUX_ORGANIZATION"),
		bucket:       os.Getenv("INFLUX_BUCKET"),
	}
}

// Address returns the InfluxDB address from the influxConfig.
func (i *influxConfig) Address() string {
	return i.address
}

// AuthToken returns the InfluxDB authorization token from the influxConfig.
func (i *influxConfig) AuthToken() string {
	return i.authToken
}

// Organization returns the InfluxDB organization from the influxConfig.
func (i *influxConfig) Organization() string {
	return i.organization
}

// Bucket returns the InfluxDB bucket from the influxConfig.
func (i *influxConfig) Bucket() string {
	return i.bucket
}
