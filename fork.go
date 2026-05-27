package immortal

// Fork crete a new process
func Fork() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Setsid is used to detach the process from the parent (normally a shell)
//
// The disowning of a child process is accomplished by executing the system call
// setpgrp() or setsid(), (both of which have the same functionality) as soon as
// the child is forked. These calls create a new process session group, make the
// child process the session leader, and set the process group ID to the process
// ID of the child. https://bsdmag.org/unix-kernel-system-calls/
