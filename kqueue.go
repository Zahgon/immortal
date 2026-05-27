//go:build (freebsd || netbsd || openbsd || dragonfly || darwin) && (amd64 || arm64)
// +build freebsd netbsd openbsd dragonfly darwin
// +build amd64 arm64

package immortal

// WatchDir check for changes on a directory via Kqueue EVFILT_VNODE
func WatchDir(dir string, ch chan<- string) error { _ = "STUB: not implemented"; return nil }

// create kevent

// wait for an event

// Move to next event

// WatchFile check for changes on a file via kqueue EVFILT_VNODE
func WatchFile(f string, ch chan<- string) error { _ = "STUB: not implemented"; return nil }

// create kevent

// wait for an event

// do something

// Move to next event
