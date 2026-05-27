//go:build (freebsd || netbsd || openbsd || darwin) && arm
// +build freebsd netbsd openbsd darwin
// +build arm

package immortal

// WatchPid check pid changes
func (d *Daemon) WatchPid(pid int, ch chan<- error) { _ = "STUB: not implemented"; return }
