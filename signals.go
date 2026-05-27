package immortal

import (
	"net/http"
)

// SignalResponse struct to return the error in json format
type SignalResponse struct {
	Err string
}

// HandleSignal send signals to the current process
func (d *Daemon) HandleSignal(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"

	// get signal from request params
	return
}

// a: Alarm. Send the service an ALRM signal.

// c: Continue. Send the service a CONT signal.

// d: Down. If the service is running, send it a TERM signal. After it stops, do not restart it.

// h: Hangup. Send the service a HUP signal.

// halt: down + exit
// A restart will only happen when using immortaldir

// i: Interrupt. Send the service an INT signal.

// in: TTIN. Send the service a TTIN signal.

// k: Kill. Send the service a KILL signal.

// o: Once. If the service is not running, start it. Do not restart it if it stops.

// ou: TTOU. Send the service a TTOU signal.

// s: stop. Send the service a STOP signal.

// q: QUIT. Send the service a QUIT signal.

// t: Terminate. Send the service a TERM signal.

// u: Up. If the service is not running, start it. If the service stops, restart it.

// 1: USR1. Send the service a USR1 signal.

// 2: USR2. Send the service a USR2 signal.

// w: WINCH. Send the service a WINCH signal.

// x: Exit. If you use this option on a stable system, you're doing something wrong.
// the supervisor is designed to run forever.

// return the error on the Response json encoded
