//go:build freebsd || netbsd || openbsd || dragonfly || darwin
// +build freebsd netbsd openbsd dragonfly darwin

package immortal

import (
	"sync"
)

// ScanDir struct
type ScanDir struct {
	scandir  string
	sdir     string
	services sync.Map
	watch    chan string
}

// NewScanDir returns ScanDir struct
func NewScanDir(path string) (*ScanDir, error) { _ = "STUB: not implemented"; return nil, nil }

// Start check for changes on directory
func (s *ScanDir) Start(ctl Control) { _ = "STUB: not implemented"; return }

// create supervise directory (/var/run/immortal) if doesn't exists
// IMMORTAL_SDIR

// check for changes on sdir helps to restart stoped services

// check for new services on scandir

// start with scandir

// Try to start services that are stopped depending on require/require_cmd

// every 5 seconds

// based on kqueue response

// RESTART, after halting a service, this will
// start stopped services after 1 second of receiving the signal

// restart if file changed

// try to start before via socket

// keep retrying

// remove service

// Scandir searches for *.yml if file changes it will reload(stop-start)
func (s *ScanDir) Scandir(ctl Control) error { _ = "STUB: not implemented"; return nil }

// start or restart if service is not in map or file lock don't exists

// Block for 100 ms on each call to kevent (WatchFile)

// WatchFile - react on file changes
func (s *ScanDir) WatchFile(path string) { _ = "STUB: not implemented"; return }

// try 3 times sleeping i*100ms between retries
