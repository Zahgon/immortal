package immortal

import (
	"time"
)

// Supervisor for the process
type Supervisor struct {
	daemon  *Daemon
	process *process
	pid     int
	wait    time.Duration
}

// Supervise keep daemon process up and running
func Supervise(d *Daemon) error {
	_ = "STUB: not implemented"
	// start a new process
	return nil
}

// Start loop forever
func (s *Supervisor) Start() error { _ = "STUB: not implemented"; return nil }

// get exit code
// TODO check EXIT from kqueue since we don't know the exit code there

// Check for post_exit command

// stop or exit based on the retries

// stop don't exit

// follow the new pid instead of trying to call run again unless the new pid dies

// ReStart create a new process
func (s *Supervisor) ReStart() { _ = "STUB: not implemented"; return }

// loop again but wait 1 seccond before trying

// Terminate handle process termination
func (s *Supervisor) Terminate(err error) bool { _ = "STUB: not implemented"; return false }

// set end time

// unlock, or lock once

// WatchPid returns EXIT

// calculate time for next reboot (avoids high CPU usage)

// behavior based on the retries

//  0 run only once (don't retry)

// +1 run N times

// -1 run forever

// FollowPid check if process still up and running if it is, follow the pid,
// monitor the existing pid created by the process instead of creating
// another process
func (s *Supervisor) FollowPid(err error) { _ = "STUB: not implemented"; return }

// check if pid in file is valid

// overwrite original (defunct) pid with the fpid in order to be available to send signals

// if cmd exits or process is kill
