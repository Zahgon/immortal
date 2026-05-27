//go:build (freebsd || netbsd || openbsd || dragonfly || darwin) && (amd64 || arm64)
// +build freebsd netbsd openbsd dragonfly darwin
// +build amd64 arm64

package immortal

// WatchPid check pid changes
func (d *Daemon) WatchPid(pid int, ch chan<- error) { _ = "STUB: not implemented"; return }

// create kevent
