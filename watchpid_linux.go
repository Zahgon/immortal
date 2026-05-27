//go:build linux
// +build linux

package immortal

// WatchPid check pid changes
func (d *Daemon) WatchPid(pid int, ch chan<- error) { _ = "STUB: not implemented"; return }
