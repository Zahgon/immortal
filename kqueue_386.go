//go:build (freebsd || netbsd || openbsd || darwin) && 386
// +build freebsd netbsd openbsd darwin
// +build 386

package immortal

// WatchDir check for changes on a directory via Kqueue EVFILT_VNODE
func WatchDir(dir string, ch chan<- string) error { _ = "STUB: not implemented"; return nil }

// create kevent

// wait for an event

// Move to next event

// WatchFile check for changes on a file via kqueue EVFILT_VNODE
func WatchFile(f string, ch chan<- string) error { _ = "STUB: not implemented"; return nil }

// NOTE_WRITE and NOTE_ATTRIB returns twice, if removing NOTE_ATTRIB (touch) will not work

// create kevent

// wait for an event

// do something

// Move to next event
