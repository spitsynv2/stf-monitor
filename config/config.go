package config

import (
	"log"
	"os"
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

	val := os.Getenv("PULL_COOLDOWN")
	duration, err := time.ParseDuration(val)
	if err != nil {
		log.Fatalf("invalid PULL_COOLDOWN: %v", err)
	}
	Conf.PullCooldown = duration
}
