package immortal

import (
	"net/http"
)

// Status struct
type Status struct {
	Pid    int    `json:"pid"`
	Up     string `json:"up,omitempty"`
	Down   string `json:"down,omitempty"`
	Cmd    string `json:"cmd"`
	Fpid   bool   `json:"fpid"`
	Count  int    `json:"count"`
	Status string `json:"status,omitempty"`
}

// Listen creates a unix socket used for control the daemon
func (d *Daemon) Listen() (err error) { _ = "STUB: not implemented"; return nil }

// close socket when process finishes (after cmd.Wait())

// HandleStatus return process status
func (d *Daemon) HandleStatus(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

//  only if process is running

// return status in json
