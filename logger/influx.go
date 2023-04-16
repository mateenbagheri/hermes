package logger

import (
	"bytes"
	"encoding/json"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/rs/zerolog"
)

// The influxWriter struct provides an implementation of the io.Writer interface
// that writes logs to InfluxDB.
type influxWriter struct {
	client influxdb2.Client
	api    api.WriteAPI
}

// Write writes the given byte slice to InfluxDB.
func (inf *influxWriter) Write(p []byte) (int, error) {
	evt, err := inf.Decode(p)
	if err != nil {
		return -1, err
	}

	point := inf.getWritePoint(evt)

	inf.api.WritePoint(point)

	go func() {
		for {
			if err := <-inf.api.Errors(); err != nil {
				panic(err)
			}
		}
	}()

	return len(p), nil
}

// Decode decodes the given byte slice into a map[string]any.
func (inf *influxWriter) Decode(p []byte) (map[string]any, error) {
	var evt map[string]any

	decoder := json.NewDecoder(bytes.NewReader(p))
	if err := decoder.Decode(&evt); err != nil {
		return nil, fmt.Errorf("cannot decode event: %s", err)
	}

	return evt, nil
}

// getWritePoint constructs a new InfluxDB point from the given map[string]any.
func (inf *influxWriter) getWritePoint(evt map[string]any) *write.Point {
	point := influxdb2.NewPointWithMeasurement("log")

	for k, v := range evt {
		switch k {
		// Ignore fields that are already present in the measurement name.
		case zerolog.TimestampFieldName:
			continue
		case zerolog.LevelFieldName:
			point = point.AddTag(zerolog.LevelFieldName, v.(string))
		case zerolog.CallerFieldName:
			point = point.AddTag(zerolog.CallerFieldName, v.(string))
		case zerolog.MessageFieldName:
			point = point.AddTag(zerolog.MessageFieldName, v.(string))
		default:
			point = point.AddField(k, v)
		}
	}

	return point
}
