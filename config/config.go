package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	StfProviders map[string]string
	SkipSerials  map[string]bool
	PullCooldown time.Duration
}

var Conf = Config{}

func init() {
	Conf.StfProviders = make(map[string]string)
	providers := strings.Split(os.Getenv("STF_PROVIDERS"), ";")
	tokens := strings.Split(os.Getenv("STF_TOKENS"), ";")

	Conf.StfProviders = make(map[string]string)

	for i := range providers {
		if i < len(tokens) {
			Conf.StfProviders[providers[i]] = tokens[i]
		}
	}

	Conf.SkipSerials = make(map[string]bool)
	skipSerials := strings.Split(os.Getenv("SKIP_SERIALS"), ";")
	for _, val := range skipSerials {
		Conf.SkipSerials[val] = true
	}

	val := os.Getenv("STF_API_REQUEST_INTERVAL")
	duration, err := time.ParseDuration(val)
	if err != nil {
		sec, err2 := strconv.Atoi(val)
		if err2 != nil {
			log.Fatalf("invalid STF_API_REQUEST_INTERVAL: %v", err)
		}
		duration = time.Duration(sec) * time.Second
	}
	Conf.PullCooldown = duration
}
