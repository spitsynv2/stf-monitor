package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spitsynv2/stf-monitor/api"
	"github.com/spitsynv2/stf-monitor/config"
	"github.com/spitsynv2/stf-monitor/model"
)

// Global: provider -> (device serial -> status)
var providerStates = make(map[string]map[string]model.DeviceStatus)

func pollDevices() {
	for provider := range config.Conf.StfProviders {
		devices, err := api.GetDevices(provider)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if _, exists := providerStates[provider]; !exists {
			providerStates[provider] = make(map[string]model.DeviceStatus)
		}

		for _, d := range devices {
			if config.Conf.SkipSerials[d.Serial] {
				continue
			}

			now := time.Now()
			current := d.Present

			prev, exists := providerStates[provider][d.Serial]
			if !exists {
				providerStates[provider][d.Serial] = model.DeviceStatus{
					Present:   current,
					ChangedAt: now,
					Duration:  0,
				}
				continue
			}

			if prev.Present != current {
				if !current {
					fmt.Printf("[%s] %s (%s) from %s went OFFLINE\n",
						now.Format(time.RFC3339), d.Serial, d.Model, provider)
				} else {
					downtime := now.Sub(prev.ChangedAt).Round(time.Second)
					fmt.Printf("[%s] %s (%s) from %s came ONLINE after %s\n",
						now.Format(time.RFC3339), d.Serial, d.Model, provider, downtime)
				}

				providerStates[provider][d.Serial] = model.DeviceStatus{
					Present:   current,
					ChangedAt: now,
					Duration:  0,
				}
			} else {
				prev.Duration = now.Sub(prev.ChangedAt).Round(time.Second)
				providerStates[provider][d.Serial] = prev
			}
		}
	}
}

func printStatusDurations() {
	fmt.Println("---- Device Status Durations ----")
	for provider, devices := range providerStates {
		fmt.Printf("Provider: %s\n", provider)
		for serial, status := range devices {
			state := "OFFLINE"
			if status.Present {
				state = "ONLINE"
			}
			fmt.Printf("  %s - %s for %s\n", serial, state, status.Duration)
		}
	}
	fmt.Println("---------------------------------")
}

func main() {
	root, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(api.TimeoutMiddleware(1 * time.Minute))
	api.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":7575",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server started on :7575")

loop:
	for {
		select {
		case <-ticker.C:
			pollDevices()
			//printStatusDurations()
			model.UpdateMetrics(providerStates)

		case <-root.Done():
			log.Println("received shutdown signal, stopping poll loop...")
			break loop
		}
	}

	log.Println("shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exiting")
}
