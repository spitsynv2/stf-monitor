package config

import (
	"os"
	"strings"
)

type Config struct {
	StfProviders map[string]string
	SkipSerials  map[string]bool
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
}
