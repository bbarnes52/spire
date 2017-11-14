package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	workload "github.com/spiffe/spire/proto/api/workload"
)

// Workload is the component that consumes Workload API and renews certs
type Workload struct {
	workloadClient        workload.WorkloadClient
	workloadClientContext context.Context
	timeout               int
}

// NewWorkload creates a new workload
func NewWorkload(workloadClientContext context.Context, workloadClient workload.WorkloadClient, timeout int) *Workload {
	return &Workload{
		workloadClientContext: workloadClientContext,
		workloadClient:        workloadClient,
		timeout:               timeout,
	}
}

// RunDaemon starts the main loop
func (w *Workload) RunDaemon() error {
	// Create channel for interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Create timer for timeout
	timeoutTimer := time.NewTimer(time.Second * time.Duration(w.timeout))

	// Main loop
	for {
		// Fetch certificates
		ttl, err := w.fetchBundles()
		if err != nil {
			return err
		}

		// Create timer for TTL
		timer := time.NewTimer(time.Second * time.Duration(ttl))

		// Wait for either timer or interrupt signal
		log("Will wait for TTL (%d seconds)\n", ttl)
		select {
		case <-timer.C:
			log("Time is up! Will renew cert.\n")
			// Continue
		case <-timeoutTimer.C:
			log("Global timeout! Will exit.\n")
			return nil
		case <-interrupt:
			log("Interrupted! Will exit.\n")
			return nil
		}
	}
}

func (w *Workload) fetchBundles() (ttl int32, err error) {
	bundles, err := w.workloadClient.FetchAllBundles(w.workloadClientContext, &workload.Empty{})
	if err != nil {
		return
	}

	if len(bundles.Bundles) == 0 {
		err = errors.New("Fetched zero bundles")
		return
	}

	ttl = bundles.Ttl
	return
}

func log(format string, a ...interface{}) {
	fmt.Print(time.Now().Format(time.Stamp), ": ")
	fmt.Printf(format, a...)
}