//go:build (freebsd || netbsd || openbsd || darwin) && 386
// +build freebsd netbsd openbsd darwin
// +build 386

package immortal

// WatchPid check pid changes
func (d *Daemon) WatchPid(pid int, ch chan<- error) { _ = "STUB: not implemented"; return }
