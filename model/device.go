package model

import (
	"encoding/json"
	"time"
)

type Device struct {
	Serial       string `json:"serial"`
	Name         string `json:"name"`
	Model        string `json:"model"`
	DeviceType   string `json:"deviceType"`
	Platform     string `json:"platform"`
	Version      string `json:"version"`
	Manufacturer string `json:"manufacturer,omitempty"`
	MarketName   string `json:"marketName,omitempty"`

	Present         bool    `json:"present"`
	Ready           bool    `json:"ready"`
	Using           bool    `json:"using"`
	Status          int     `json:"status"`
	StatusChangedAt string  `json:"statusChangedAt"`
	Usage           *string `json:"usage,omitempty"`
	UsageChangedAt  string  `json:"usageChangedAt"`

	Battery struct {
		Level  int             `json:"level"`
		Status string          `json:"status"`
		Health string          `json:"health"`
		Temp   json.RawMessage `json:"temp"`
		Source string          `json:"source"`
	} `json:"battery"`

	Display struct {
		Width   int     `json:"width"`
		Height  int     `json:"height"`
		Density float64 `json:"density"`
		FPS     float64 `json:"fps"`
		URL     string  `json:"url"`
	} `json:"display"`

	Network *struct {
		Connected bool   `json:"connected"`
		Type      string `json:"type"`
		Roaming   bool   `json:"roaming"`
	} `json:"network,omitempty"`

	Provider struct {
		Name string `json:"name"`
	} `json:"provider"`
	Group struct {
		Name string `json:"name"`
	} `json:"group"`
	Owner *struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"owner,omitempty"`

	CreatedAt         string `json:"createdAt"`
	PresenceChangedAt string `json:"presenceChangedAt"`
}

type DeviceResponse struct {
	Success     bool     `json:"success"`
	Description string   `json:"description"`
	Devices     []Device `json:"devices"`
}

type DeviceStatus struct {
	Present   bool
	ChangedAt time.Time
	Duration  time.Duration
}
