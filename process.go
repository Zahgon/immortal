package immortal

import (
	"os"
	"os/exec"
	"syscall"
	"time"
)

// Process interface
type Process interface {
	Kill() error
	Pid() int
	Signal(sig syscall.Signal) error
	Start() (*process, error)
	GetProcess() *process
}

type process struct {
	*Config
	Logger
	LoggerStderr Logger
	cmd          *exec.Cmd
	errch        chan error
	quit         chan struct{}
	sTime, eTime time.Time
}

// SetEnv set environment variables - If the Cmd.Env contains duplicate
// environment keys, only the last value in the slice for each duplicate
// key is used.
func (p *process) SetEnv(env []string) { _ = "STUB: not implemented"; return }

// SetsysProcAttr - set process group ID and owner (run on behalf)
func (p *process) SetsysProcAttr() error { _ = "STUB: not implemented"; return nil }

// Set process group ID to Pgid, or, if Pgid == 0, to new pid.
// Child's process group ID if Setpgid.

// set owner

// set the attributes

// Start runs the command
func (p *process) Start() (*process, error) {
	_ = "STUB: not implemented"
	// command obtained from Config parent
	return nil, nil
}

// change working directory

// set environment variables

// set sysProcAttr

// log only if are available loggers

// create the pipes for Stdout

// create the pipes for Stdout

// create the pipes for Stdout

// Start the process

// set start time

// wait process to finish in a goroutine

// Wait - wait process to finish
func (p *process) Wait(stdout, stderr *os.File) { _ = "STUB: not implemented"; return }

// Kill the entire Process group.
func (p *process) Kill() error { _ = "STUB: not implemented"; return nil }

// Pid return Process PID
func (p *process) Pid() int { _ = "STUB: not implemented"; return 0 }

// Signal sends a signal to the Process
func (p *process) Signal(sig syscall.Signal) error { _ = "STUB: not implemented"; return nil }

// GetProccess
func (p *process) GetProcess() *process {
	_ = "STUB: not implemented"

	// NewProcess return process instance
	return nil
}

func NewProcess(cfg *Config) *process { _ = "STUB: not implemented"; return nil }
