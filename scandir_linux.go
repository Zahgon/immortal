//go:build linux
// +build linux

package immortal

import (
	"time"
)

// ScanDir struct
type ScanDir struct {
	scandir       string
	sdir          string
	services      map[string]string
	timeMultipler time.Duration
}

// NewScanDir returns ScanDir struct
func NewScanDir(path string) (*ScanDir, error) { _ = "STUB: not implemented"; return nil, nil }

// Start scans directory every 5 seconds
func (s *ScanDir) Start(ctl Control) { _ = "STUB: not implemented"; return }

// Scanner searches for run.yml if file changes it will reload(stop-start)
func (s *ScanDir) Scanner(ctl Control) {
	_ = "STUB: not implemented"
	// var services used to keep track of what services should be removed if they don't
	// exist any more
	return
}

// only use .yml files

// add service to services map or reload if file has been changed

// check if file hasn't been changed since last tick (5 seconds)

// restart = term + start

// try to start before via socket

// keep retrying

// find for .yml files

// halts services that don't exist anymore
