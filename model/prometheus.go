package model

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	devicePresent = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "device_present",
			Help: "Whether the device is present (1) or not (0).",
		},
		[]string{"provider", "serial"},
	)

	deviceChangedAt = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "device_changed_at",
			Help: "Timestamp when the device status last changed (Unix seconds).",
		},
		[]string{"provider", "serial"},
	)

	deviceDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "device_duration_seconds",
			Help: "How long the device has been in the current state (seconds).",
		},
		[]string{"provider", "serial"},
	)
)

func init() {
	prometheus.MustRegister(devicePresent, deviceChangedAt, deviceDuration)
}

func UpdateMetrics(providerStates map[string]map[string]DeviceStatus) {
	for provider, devices := range providerStates {
		for serial, status := range devices {
			if status.Present {
				devicePresent.WithLabelValues(provider, serial).Set(1)
			} else {
				devicePresent.WithLabelValues(provider, serial).Set(0)
			}

			deviceChangedAt.WithLabelValues(provider, serial).
				Set(float64(status.ChangedAt.Unix()))

			deviceDuration.WithLabelValues(provider, serial).
				Set(status.Duration.Seconds())
		}
	}
}
