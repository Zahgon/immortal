package immortal

import (
	"sync"
	"time"
)

// Daemon struct
type Daemon struct {
	sync.RWMutex
	cfg            *Config
	count          int
	fpid           bool
	lock, lockOnce uint32
	process        *process
	quit, run      chan struct{}
	sTime          time.Time
	supDir         string
	wg             sync.WaitGroup
}

// Run returns a process instance
func (d *Daemon) Run(p Process) (*process, error) { _ = "STUB: not implemented"; return nil, nil }

// return if process is running

// increment count by 1

// to print remaininig seconds to start cmd == nil

// write parent pid

// write child pid

// not following a pid

// WritePid write pid to file
func (d *Daemon) WritePid(file string, pid int) error { _ = "STUB: not implemented"; return nil }

// IsRunning check if process is running
func (d *Daemon) IsRunning(pid int) bool { _ = "STUB: not implemented"; return false }

// ReadPidFile read pid from file if error returns pid 0
func (d *Daemon) ReadPidFile(pidfile string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// New creates a new daemon
func New(cfg *Config) (*Daemon, error) {
	_ = "STUB: not implemented"

	// create supervise directory in specified directory
	// defaults to /var/run/immortal/<app>
	return nil, nil
}

// create an .immortal dir on $HOME user when calling immortal directly
// and not using immortal-dir, this helps to run immortal-ctl and
// check status of all daemons

// create supervise dir

// lock

// resource temporarily unavailable

// remove previous socket in case exists
